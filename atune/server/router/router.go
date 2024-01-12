package router

import (
	"net/http"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/atune-plugin/config"
	"openeuler.org/PilotGo/atune-plugin/controller"
	"openeuler.org/PilotGo/atune-plugin/plugin"
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
	plugin.GlobalClient.RegisterHandlers(router)

	atune := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)
	{
		atune.POST("atune_install", controller.AtuneClientInstall)
		atune.POST("atune_uninstall", controller.AtuneClientRemove)
		atune.GET("all", controller.GetAtuneAll)
		atune.GET("info", controller.GetAtuneInfo)
	}

	dbtune := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)
	{
		dbtune.GET("tunes", controller.QueryTunes)
		dbtune.POST("save_tune", controller.SaveTune)
		dbtune.POST("update_tune", controller.UpdateTune)
		dbtune.DELETE("delete_tune", controller.DeleteTune)
		dbtune.GET("search_tune", controller.SearchTune)
	}

	task := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)
	{
		task.POST("task_new", controller.CreatTask)
		task.GET("tasks", controller.TaskLists)
		task.DELETE("task_delete", controller.DeleteTask)
		task.GET("task_search", controller.SearchTask)
	}

	restune := router.Group("/plugin/" + plugin.GlobalClient.PluginInfo.Name)
	{
		restune.DELETE("result_delete", controller.DeleteResult)
		restune.GET("result_search", controller.SearchResult)
	}
}

func StaticRouter(router *gin.Engine) {
	router.Static("/plugin/atune/static", "../web/dist/static")
	router.StaticFile("/plugin/atune", "../web/dist/index.html")

	// 解决页面刷新404的问题
	router.NoRoute(func(c *gin.Context) {
		logger.Error("process noroute: %s", c.Request.URL)
		if !strings.HasPrefix(c.Request.RequestURI, "/plugin/atune/*path") {
			c.File("../web/dist/index.html")
			return
		}
		c.Status(http.StatusNotFound)
	})
}
