package provider

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var appPort string

type AppOptions struct {
	ApiPort           string
	RouterConstructor interface{}
	AppConstructors   []interface{}
	ResponseWriter    http.ResponseWriter
	Request           *http.Request
}

type Api interface {
	Run()
	Routes() *gin.Engine
}

func NewProductionApp(options AppOptions) *fx.App {
	appPort = options.ApiPort
	return fx.New(
		fx.Provide(fx.Annotate(options.RouterConstructor, fx.As(new(Api)))),
		fx.Provide(options.AppConstructors...),
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

func NewDevelopmentApp(options AppOptions) *fx.App {
	appPort = options.ApiPort
	return fx.New(
		fx.Provide(fx.Annotate(options.RouterConstructor, fx.As(new(Api)))),
		fx.Provide(options.AppConstructors...),
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
