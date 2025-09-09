package router

import (
	"net/http"
	"strings"

	"gitee.com/openeuler/PilotGo/sdk/logger"
	"github.com/gin-gonic/gin"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/global"
	customscripts "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/custom_scripts"
	dangerousrule "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/dangerous_rule"
	scriptlibrary "openeuler.org/PilotGo/PilotGo-plugin-automation/internal/module/script_library"
	"openeuler.org/PilotGo/PilotGo-plugin-automation/internal/service"
)

func HttpServerInit() *gin.Engine {
	server := initRouters()
	frontendStaticRouter(server)
	clientResister(server)

	logger.Debug("http server successfully started, listening on %s", global.App.HttpAddr)
	return server
}

// 后端路由
func initRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(logger.RequestLogger([]string{}))

	// 注册各自的路由模块
	api := Router.Group("/plugin/automation")
	customscripts.CustomScriptsHandler(api)
	scriptlibrary.ScriptLibraryHandler(api)
	dangerousrule.DangerousRuleHandler(api)
	return Router
}

// 加载前端静态资源
func frontendStaticRouter(router *gin.Engine) {
	router.Static("/plugin/automation/assets", "../web/dist/assets")
	router.StaticFile("/plugin/automation", "../web/dist/index.html")

	// 解决页面刷新404的问题
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/plugin/automation/api") {
			c.File("../web/dist/index.html")
			return
		}
		c.AbortWithStatus(http.StatusNotFound)
	})
}

func clientResister(router *gin.Engine) {
	global.App.Client.RegisterHandlers(router)
	service.GetTags()
}
