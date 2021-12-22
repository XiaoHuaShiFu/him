package group

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"github.com/jiaxwu/him/common/jsons"
	"github.com/jiaxwu/him/config/log"
	"os"
	"time"
)

// GroupQRCode 群二维码
type GroupQRCode struct {
	GroupID        uint64 `json:"GroupID"`        // 群编号
	InviterID      uint64 `json:"InviterID"`      // 邀请者编号
	ExpirationTime int64  `json:"ExpirationTime"` // 过期时间
}

const (
	GroupQRCodeEffectiveTime = time.Hour * 24 * 7                  // 群二维码有效期
	GroupQRCodePrefix        = "https://www.him.com/group/qrcode/" // 群二维码前缀
)

const (
	groupQRCodePrivateKeyFilePath = "service/group/group_qrcode_private_key.pem" // 私钥文件路径
	groupQRCodePublicKeyFilePath  = "service/group/group_qrcode_public_key.pem"  // 公钥文件路径
)

var (
	// 群二维码私钥
	groupQrcodePrivateKey = mustGroupQRCodePrivateKey(groupQRCodePrivateKeyFilePath)
	// 群二维码公钥
	groupQrcodePublicKey = mustGroupQRCodePublicKey(groupQRCodePublicKeyFilePath)
)

// 群二维码解密（使用私钥）
func groupQRCodeDecrypt(base64GroupQRCode string) (*GroupQRCode, error) {
	// base64解码
	encryptedGroupQRCode, err := base64.StdEncoding.DecodeString(base64GroupQRCode)
	if err != nil {
		return nil, err
	}
	// rsa解密
	bytesGroupQRCode, err := rsa.DecryptPKCS1v15(rand.Reader, groupQrcodePrivateKey, encryptedGroupQRCode)
	if err != nil {
		return nil, err
	}
	// 去除前缀
	bytesGroupQRCode = bytesGroupQRCode[len(GroupQRCodePrefix):]
	// json解码
	var groupQRCode GroupQRCode
	if err := json.Unmarshal(bytesGroupQRCode, &groupQRCode); err != nil {
		return nil, err
	}
	return &groupQRCode, nil
}

// 群二维码加密（使用公钥）
func groupQRCodeEncrypt(groupQRCode *GroupQRCode) (string, error) {
	// 加上前缀
	bytesGroupQRCode := append([]byte(GroupQRCodePrefix), jsons.MarshalToBytes(groupQRCode)...)
	// rsa加密
	encryptedGroupQRCode, err := rsa.EncryptPKCS1v15(rand.Reader, groupQrcodePublicKey, bytesGroupQRCode)
	if err != nil {
		return "", err
	}
	// base64编码
	return base64.StdEncoding.EncodeToString(encryptedGroupQRCode), nil
}

// 群二维码公钥
func mustGroupQRCodePublicKey(path string) *rsa.PublicKey {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.WithError(err).Fatal()
	}
	block, _ := pem.Decode(buf)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.WithError(err).Fatal("parse exception")
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.Fatal("can not converse the type of pub to *rsa.PublicKey")
	}
	return publicKey
}

// 群二维码私钥
func mustGroupQRCodePrivateKey(path string) *rsa.PrivateKey {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.WithError(err).Fatal()
	}
	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.WithError(err).Fatal("parse exception")
	}
	return privateKey
}
