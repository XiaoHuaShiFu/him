package main

import (
	"go.uber.org/fx"
	"him/conf"
	"him/conf/db"
	"him/conf/logger"
	"him/conf/rdb"
	"him/conf/validate"
	"him/service/server"
	"him/service/service/auth"
	authHandler "him/service/service/auth/handler"
	"him/service/service/msg/gateway"
	"him/service/service/sm"
	"him/service/service/user/profile"
	"him/service/wrap"
)

func main() {
	NewApp().Run()
}

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLogger,
			validate.NewValidate,
			conf.NewConf,
			db.NewDB,
			rdb.NewRDB,
			server.NewEngine,
			server.NewServer,
			wrap.NewWrapper,
			sm.NewService,
			fx.Annotate(
				auth.NewAuthEventProducer,
				fx.ResultTags(`name:"AuthEventProducer"`),
			),
			fx.Annotate(
				auth.NewService,
				fx.ParamTags(`name:"AuthEventProducer"`),
			),
			fx.Annotate(
				profile.NewUserAvatarBucketOSSClient,
				fx.ResultTags(`name:"UserAvatarBucketOSSClient"`),
			),
			fx.Annotate(
				profile.NewUserProfileEventProducer,
				fx.ResultTags(`name:"UserProfileEventProducer"`),
			),
			fx.Annotate(
				profile.NewService,
				fx.ParamTags(`name:"UserAvatarBucketOSSClient"`, `name:"UserProfileEventProducer"`),
			),
		),
		fx.Invoke(
			authHandler.RegisterHandler,
			profile.RegisterUserProfileHandler,
			gateway.NewGatewayHandler,
			profile.NewAuthEventConsumer,
			fx.Annotate(
				profile.NewUserAvatarClearTask,
				fx.ParamTags(`name:"UserAvatarBucketOSSClient"`),
			),
			server.Start,
		),
		fx.WithLogger(logger.NewFxEventLogger),
	)
}
