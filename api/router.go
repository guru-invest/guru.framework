package api

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MethodTypes int

const (
	GET 	 MethodTypes = iota
	POST
	PUT
	DELETE
	OPTIONS
)

type routes []mapper
var _routes = routes{}
var _staticRoutes []staticMapper
var srv *http.Server

type mapper struct{
	Method MethodTypes
	Pattern string
	Handler gin.HandlerFunc
}

type staticMapper struct{
	Pattern string
	Directory string
}

func addRoutesToMapper(router gin.IRouter){
	for _, route := range _routes  {
		switch route.Method {
		case POST:
			router.POST(route.Pattern, route.Handler)
		case PUT:
			router.PUT(route.Pattern, route.Handler)
		case DELETE:
			router.DELETE(route.Pattern, route.Handler)
		case OPTIONS:
			router.OPTIONS(route.Pattern, route.Handler)
		default:
			router.GET(route.Pattern, route.Handler)
		}
	}
}


func addStaticRoutesToMapper(router gin.IRouter){
	if len(_staticRoutes) >0{
		for _,route := range _staticRoutes{
			router.StaticFS(route.Pattern, http.Dir(route.Directory))
		}
	}
}

func Static(pattern string, directory string){
	_route := staticMapper{
		Pattern: pattern,
		Directory: directory,
	}
	_staticRoutes = append(_staticRoutes, _route)
}


func AddRoute(method MethodTypes, pattern string, handler gin.HandlerFunc){
	_route := mapper{
		method,
		"/"+pattern,
		handler,
	}
	_routes = append(_routes, _route)
}

func createEngine(apiVersion string) *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	version := router.Group("/" + apiVersion)
	{
		addRoutesToMapper(version)
		addStaticRoutesToMapper(version)
	}
	return router
}

func InitRoutering(port string, apiVersion string, isServerless bool) {
	router := createEngine(apiVersion)
	if isServerless {
		gateway.ListenAndServe(":" + port, router)
	}else{
		srv = &http.Server{
			Addr:    ":"+port,
			Handler: router,
		}
		srv.ListenAndServe()
	}
}