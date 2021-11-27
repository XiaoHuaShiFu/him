package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	pb "github.com/xiaohuashifu/him/api/authnz/authz"
	"github.com/xiaohuashifu/him/conf"
	"github.com/xiaohuashifu/him/service/common"
	"github.com/xiaohuashifu/him/service/mq"
	"github.com/xiaohuashifu/him/service/service/authnz/authz/code"
	"github.com/xiaohuashifu/him/service/service/authnz/authz/db"
	"github.com/xiaohuashifu/him/service/service/authnz/authz/model"
	smsModel "github.com/xiaohuashifu/him/service/service/sms/model"
	smsService "github.com/xiaohuashifu/him/service/service/sms/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthzService struct {
	db                 *gorm.DB
	rdb                *redis.Client
	validate           *validator.Validate
	logger             *logrus.Logger
	smsService         *smsService.SMSService
	loginEventProducer rocketmq.Producer
	config             *conf.Config
}

func NewAuthzService(loginEventProducer rocketmq.Producer, db *gorm.DB, rdb *redis.Client, validate *validator.Validate,
	logger *logrus.Logger, smsService *smsService.SMSService, config *conf.Config) *AuthzService {
	return &AuthzService{
		db:                 db,
		rdb:                rdb,
		validate:           validate,
		logger:             logger,
		smsService:         smsService,
		config:             config,
		loginEventProducer: loginEventProducer,
	}
}

// Login 登录
func (s *AuthzService) Login(req *pb.LoginReq) (*pb.LoginRsp, *common.Error) {
	var (
		rsp *pb.LoginRsp
		err *common.Error
	)
	if req.Type == pb.LoginType_LoginTypeSMVerCode {
		// sms登录
		rsp, err = s.loginBySMS(req.Phone, req.AuthCode)
	} else if req.Type == pb.LoginType_LoginTypePwd {
		// pwd登录
		rsp, err = s.loginByPwd(req.Username, req.Password)
	} else {
		return nil, common.NewError(common.ErrCodeInvalidParameter)
	}
	if err != nil {
		return nil, err
	}

	// 发送登录事件
	s.sendLoginEvent(&mq.LoginEvent{
		UserID:    rsp.UserID,
		LoginTime: uint64(time.Now().Unix()),
		LoginType: req.Type,
	})

	return rsp, nil
}

// 短信验证码登录
func (s *AuthzService) loginBySMS(phone, authCode string) (*model.LoginRsp, common.Error) {
	// 验证手机号码和验证码格式
	var req = struct {
		Phone    string `validate:"required,phone"`
		AuthCode string `validate:"required,len=6,numeric"`
	}{
		Phone:    phone,
		AuthCode: authCode,
	}
	if err := s.validate.Struct(req); err != nil {
		return nil, common.NewError(common.ErrCodeInvalidParameter)
	}

	// 验证短信验证码
	authCodeKey := s.smsAuthCodeRedisKeyForLogin(phone)
	authCodeByRedis, err := s.rdb.Get(context.Background(), authCodeKey).Result()
	if err == redis.Nil {
		return nil, common.NewError(code.InvalidParameterLoginSMSAuthCodeNotExist)
	}
	if err != nil {
		s.logger.WithField("err", err).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}
	if authCodeByRedis != authCode {
		return nil, common.NewError(code.InvalidParameterLoginSMSAuthCodeError)
	}

	// 验证成功要把验证码删了
	if err := s.rdb.Del(context.Background(), authCodeKey).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	return s.loginByPhone(phone)
}

// loginByPhone 通过手机登录
func (s AuthzService) loginByPhone(phone string) (*model.LoginRsp, common.Error) {
	// 获取用户编号
	phoneLogin := db.PhoneLogin{
		Phone: phone,
	}
	err := s.db.Where(&phoneLogin).Take(&phoneLogin).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.WithField("err", err).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}
	userID := phoneLogin.UserID

	// 如果用户还没有使用手机注册，则进行注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if userID, err = s.register(phone); err != nil {
			return nil, err.(common.Error)
		}
	}

	return s.login(userID)
}

// 密码登录
func (s *AuthzService) loginByPwd(username, password string) (*model.LoginRsp, common.Error) {
	// 获取对应的用户
	phoneLogin := db.PhoneLogin{
		Phone: username,
	}
	err := s.db.Where(&phoneLogin).Take(&phoneLogin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.NewError(code.InvalidParameterLoginUsernameOrPassword)
	}
	if err != nil {
		s.logger.WithField("err", err).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}

	// 获取密码
	passwordLogin := db.PasswordLogin{
		UserID: phoneLogin.UserID,
	}
	err = s.db.Where(&passwordLogin).Take(&passwordLogin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, common.NewError(code.InvalidParameterLoginUsernameOrPassword)
	}
	if err != nil {
		s.logger.WithField("err", err).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(passwordLogin.Password), []byte(password)); err != nil {
		return nil, common.NewError(code.InvalidParameterLoginUsernameOrPassword)
	}

	return s.login(phoneLogin.UserID)
}

// login 登录
func (s *AuthzService) login(userID uint64) (*model.LoginRsp, common.Error) {
	// 获取用户类型
	var user db.User
	if err := s.db.Take(&user).Error; err != nil {
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}

	// 检查该用户在该终端是否已经登录了
	antiKey := s.antiTokenRedisKey(userID)
	oldToken, err := s.rdb.Get(context.Background(), antiKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}
	// 如果已经存在，把原来的token过期了
	if err == nil {
		// 删除token和反向token
		tokenKey := s.tokenRedisKey(oldToken)
		if err := s.rdb.Del(context.Background(), tokenKey, antiKey).Err(); err != nil {
			s.logger.WithField("err", err).Error("rdb exception")
			return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
		}
	}

	// 生成Token并添加到Redis
	token := gofakeit.UUID()
	tokenKey := s.tokenRedisKey(token)
	session := &common.Session{
		UserID:   userID,
		UserType: common.UserType(user.Type),
	}
	sessionBytes, _ := json.Marshal(session)
	// 先插入反向Token
	if err := s.rdb.Set(context.Background(), antiKey, token, model.TokenExp).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}
	// 再插入正向token
	if err := s.rdb.Set(context.Background(), tokenKey, sessionBytes, model.TokenExp).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	return &model.LoginRsp{
		Token:  token,
		UserID: userID,
	}, nil
}

// SendSMSForLogin 发送登录需要的短信验证码
func (s *AuthzService) SendSMSForLogin(req *model.SendSMSForLoginReq) (
	*model.SendSMSForLoginRsp, common.Error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, common.NewError(common.ErrCodeInvalidParameter)
	}

	// 把验证码加到缓存
	authCode := gofakeit.DigitN(model.SMSAuthCodeLen)
	key := s.smsAuthCodeRedisKeyForLogin(req.Phone)
	expireTime := model.SMSAuthCodeExpMinute * time.Minute
	if err := s.rdb.Set(context.Background(), key, authCode, expireTime).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	// 发送验证码
	if _, err := s.smsService.SendAuthCodeForLogin(&smsModel.SendAuthCodeForLoginReq{
		Phone:     req.Phone,
		AuthCode:  authCode,
		ExpMinute: model.SMSAuthCodeExpMinute,
	}); err != nil {
		return nil, err
	}

	return &model.SendSMSForLoginRsp{}, nil
}

// 注册账号
func (s *AuthzService) register(phone string) (uint64, common.Error) {
	var user = db.User{
		Type:         uint8(common.UserTypeUser),
		RegisteredAt: uint64(time.Now().Unix()),
	}
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		// 创建手机登录
		if err := tx.Create(&db.PhoneLogin{
			UserID: user.ID,
			Phone:  phone,
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		s.logger.WithField("err", err).Error("db exception")
		return 0, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}
	return user.ID, nil
}

// BindPasswordLogin 绑定密码
func (s *AuthzService) BindPasswordLogin(req *model.BindPasswordLoginReq) (
	*model.BindPasswordLoginRsp, common.Error) {
	// 检查密码强度
	if err := s.checkPasswordLevel(req.Password); err != nil {
		return nil, err
	}

	// 获取原始密码
	passwordLogin := db.PasswordLogin{
		UserID: req.UserID,
	}
	takeErr := s.db.Where(&passwordLogin).Take(&passwordLogin).Error
	if takeErr != nil && !errors.Is(takeErr, gorm.ErrRecordNotFound) {
		s.logger.WithField("err", takeErr).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, takeErr)
	}

	// 设置新密码
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithField("err", err).Error("bcrypt GenerateFromPassword exception")
		return nil, common.WrapError(common.ErrCodeInternalError, err)
	}
	passwordLogin.Password = string(passwordBytes)

	// 如果没有绑定过密码，则创建
	if errors.Is(takeErr, gorm.ErrRecordNotFound) {
		if err := s.db.Create(&passwordLogin).Error; err != nil {
			s.logger.WithField("err", err).Error("db exception")
			return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
		}
		return &model.BindPasswordLoginRsp{}, nil
	}

	// 否则更新密码
	if err := s.db.Updates(&passwordLogin).Error; err != nil {
		s.logger.WithField("err", err).Error("db exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorDB, err)
	}
	return &model.BindPasswordLoginRsp{}, nil
}

// checkPasswordLevel 检查密码强度
func (s *AuthzService) checkPasswordLevel(password string) common.Error {
	if len(password) < 8 || len(password) > 20 {
		return common.NewError(code.InvalidParameterLoginPasswordNotMeetRequirements)
	}

	var count uint
	for _, regexp := range model.PasswordCharRegexpSet {
		if regexp.MatchString(password) {
			count++
		}
	}
	if count < 3 {
		return common.NewError(code.InvalidParameterLoginPasswordNotMeetRequirements)
	}
	return nil
}

// Logout 退出登录
func (s *AuthzService) Logout(req *model.LogoutReq) (*model.LogoutRsp, common.Error) {
	// 退出登录
	tokenKey := s.tokenRedisKey(req.Token)
	antiKey := s.antiTokenRedisKey(req.UserID)
	sha1, err := s.rdb.ScriptLoad(context.Background(), model.LogoutRedisScript).Result()
	if err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}
	if err := s.rdb.EvalSha(context.Background(), sha1, []string{tokenKey, antiKey}).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	// 发送退出登录事件
	s.sendLogoutEvent(&mq.LogoutEvent{
		UserID:     req.UserID,
		LogoutTime: uint64(time.Now().Unix()),
	})

	return &model.LogoutRsp{}, nil
}

// Authorize 权限验证
func (s *AuthzService) Authorize(req *model.AuthorizeReq) (*model.AuthorizeRsp, common.Error) {
	// 获取Session
	rsp, err := s.GetSession(&model.GetSessionReq{Token: req.Token})
	if err != nil {
		return nil, err
	}

	// 判断是否有需要的角色
	if req.UserTypes != nil && len(req.UserTypes) != 0 {
		hasRole := false
		for _, userType := range req.UserTypes {
			if rsp.Session.UserType == userType {
				hasRole = true
				break
			}
		}
		if !hasRole {
			return nil, common.NewError(common.ErrCodeForbidden)
		}
	}

	return &model.AuthorizeRsp{
		Session: rsp.Session,
	}, nil
}

// GetSession 获取Session
func (s *AuthzService) GetSession(req *model.GetSessionReq) (*model.GetSessionRsp, common.Error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, common.WrapError(code.UnauthorizedInvalidToken, err)
	}

	// 判断Token是否过期
	key := s.tokenRedisKey(req.Token)
	sessionBytes, err := s.rdb.Get(context.Background(), key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, common.NewError(code.UnauthorizedInvalidToken)
	}
	if err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	// 解析Session
	var session common.Session
	_ = json.Unmarshal(sessionBytes, &session)

	// 延长Token的过期时间
	if err := s.rdb.Expire(context.Background(), key, model.TokenExp).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	// 延长AntiToken过期时间
	antiKey := s.antiTokenRedisKey(session.UserID)
	if err := s.rdb.Expire(context.Background(), antiKey, model.TokenExp).Err(); err != nil {
		s.logger.WithField("err", err).Error("rdb exception")
		return nil, common.WrapError(common.ErrCodeInternalErrorRDB, err)
	}

	return &model.GetSessionRsp{
		Session: &session,
	}, nil
}

// smsAuthCodeRedisKeyForLogin 登录短信验证码 Redis Key
func (s *AuthzService) smsAuthCodeRedisKeyForLogin(phone string) string {
	return fmt.Sprintf("authnz:sms:authCode:%s", phone)
}

// tokenRedisKey Token Redis Key
func (s *AuthzService) tokenRedisKey(token string) string {
	return fmt.Sprintf("authnz:token:%s", token)
}

// antiTokenRedisKey 反向Token Redis Key
func (s *AuthzService) antiTokenRedisKey(userID uint64) string {
	return fmt.Sprintf("authnz:token:anti:%d", userID)
}

// sendLoginEvent 发送登录事件
func (s *AuthzService) sendLoginEvent(loginEvent *mq.LoginEvent) {
	body, _ := json.Marshal(loginEvent)
	message := primitive.NewMessage(s.config.RocketMQ.Topic, body).WithTag(string(mq.TagLoginEvent))
	s.sendEventMessage(message)
}

// sendLogoutEvent 发送退出登录事件
func (s *AuthzService) sendLogoutEvent(logoutEvent *mq.LogoutEvent) {
	body, _ := json.Marshal(logoutEvent)
	message := primitive.NewMessage(s.config.RocketMQ.Topic, body).WithTag(string(mq.TagLogoutEvent))
	s.sendEventMessage(message)
}

// sendEventMessage 发送事件消息
func (s *AuthzService) sendEventMessage(message *primitive.Message) {
	resCB := func(ctx context.Context, result *primitive.SendResult, err error) {
		s.logger.WithField("res", result).Info("send im success")
	}
	if err := s.loginEventProducer.SendAsync(context.Background(), resCB, message); err != nil {
		s.logger.WithField("err", err).Error("consumer im exception")
	}
}