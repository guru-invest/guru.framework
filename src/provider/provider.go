package provider

import (
	"context"
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"strconv"
)

var appPort string

type Api interface {
	Run()
	Routes() *gin.Engine
}

func NewProductionApp(port string, routerConstructor interface{}, appConstructors ...interface{}) *fx.App {
	appPort = port
	return fx.New(
		fx.Provide(routerConstructor, dig.As(new(Api))),
		fx.Provide(appConstructors...),
		fx.Invoke(func(lifecycle fx.Lifecycle, api Api) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(context.Context) error {
						port, _ := strconv.Atoi(appPort)
						go gateway.ListenAndServe(fmt.Sprintf(":%d", port), api.Routes())
						return nil
					},
				},
			)
		}),
		// 	TODO: LOG ON CLOUD WATCH
		//fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		//	return &fxevent.ZapLogger{Logger: logger}
		//}),
	)
}

func NewDevelopmentApp(port string, routerConstructor interface{}, appConstructors ...interface{}) *fx.App {
	appPort = port
	return fx.New(
		fx.Supply(new(Api)),
		fx.Provide(appConstructors...),
		fx.Invoke(func(lifecycle fx.Lifecycle, api Api) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(context.Context) error {
						go api.Run()
						return nil
					},
				},
			)
		}),
		// TODO: LOG ON CONSOLE AND CLOUDWATCH
		//fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		//	return &fxevent.ZapLogger{Logger: logger}
		//}),
	)
}
