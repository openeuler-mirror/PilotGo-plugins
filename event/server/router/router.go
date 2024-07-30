package router

import (
	"net/http"
	"strings"

	"gitee.com/openeuler/PilotGo-plugins/event/sdk"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
	plugin_manage "openeuler.org/PilotGo/PilotGo-plugin-event/client"
	"openeuler.org/PilotGo/PilotGo-plugin-event/config"
	"openeuler.org/PilotGo/PilotGo-plugin-event/controller"
)

func HttpServerInit(conf *config.HttpServer) error {

	go func() {
		r := setupRouter()

		logger.Info("start http service on: http://%s", conf.Addr)
		if err := r.Run(conf.Addr); err != nil {
			logger.Error("start http server failed:%v", err)
		}

	}()

	return nil
}
func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logger.RequestLogger())
	router.Use(gin.Recovery())

	registerAPIs(router)
	StaticRouter(router)

	return router
}
func registerAPIs(router *gin.Engine) {
	logger.Debug("router register")
	plugin_manage.EventClient.RegisterHandlers(router)
	sdk.RegisterEventHandlers(router, plugin_manage.EventClient)
	api := router.Group("/plugin/" + plugin_manage.EventClient.PluginInfo.Name)

	eventpublish := api.Group("")
	{
		eventpublish.PUT("publishEvent", controller.PublishEventHandler)
	}

	listener := api.Group("listener")
	{
		listener.PUT("register", controller.RegisterListenerHandler)
		listener.DELETE("unregister", controller.UnregisterListenerHandler)
		listener.DELETE("unpluginRegister", controller.UnPliginRegisterListenerHandler)
	}
}

func StaticRouter(router *gin.Engine) {
	router.Static("/plugin/event/static", "../web/dist/static")
	router.StaticFile("/plugin/event", "../web/dist/index.html")

	// 解决页面刷新404的问题
	router.NoRoute(func(c *gin.Context) {
		logger.Debug("process noroute: %s", c.Request.URL)
		if !strings.HasPrefix(c.Request.RequestURI, "/plugin/event/*path") {
			c.File("../web/dist/index.html")
			return
		}
		c.Status(http.StatusNotFound)
	})
}
