package api

import (
	"context"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	router := gin.New()
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

func RestartRouter() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Restarting...")
	srv.ListenAndServe()
}