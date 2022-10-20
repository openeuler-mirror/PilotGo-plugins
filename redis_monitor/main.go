package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	plugin_sdk "github.com/liweifeng1/plugin-sdk"
	"io/fs"
	"net/http"
	"os"
)

//go:embed dist
var frontend embed.FS

type RedisMonitorPlugin struct {
	logger hclog.Logger
	conn   *RedisService
}

func (g *RedisMonitorPlugin) OnLoad() error {
	g.logger.Info("Loading redis monitor plugin!")
	var err error
	g.conn, err = connectToRedis()
	if err != nil {
		g.logger.Error(err.Error())
		return err
	}
	go startWebServer(g)
	return nil
}

func (g *RedisMonitorPlugin) GetManifest() plugin_sdk.PluginManifest {
	return plugin_sdk.PluginManifest{
		Id:      "redisMonitor",
		Name:    "redis-monitor",
		Author:  "yibin",
		Version: "v0.0.1",
	}
}

func (g *RedisMonitorPlugin) GetConfiguration() []plugin_sdk.PluginConfig {
	return []plugin_sdk.PluginConfig{
		plugin_sdk.PluginConfig{
			Title:       "Server Port",
			Description: "Server Port",
			Key:         "server-port",
			Type:        plugin_sdk.PortValue,
			Values:      "25564",
		},
		plugin_sdk.PluginConfig{
			Title:       "Server Url",
			Description: "Server Url",
			Key:         "server-url",
			Type:        plugin_sdk.UrlValue,
			Values:      "http://localhost",
		},
	}
}

func (g *RedisMonitorPlugin) GetWebExtension() []plugin_sdk.WebExtension {
	return []plugin_sdk.WebExtension{
		{
			Type:           plugin_sdk.HTML,
			PathMatchRegex: "/",
			Source:         "main",
		},
	}
}

func (g *RedisMonitorPlugin) OnClose() error {
	g.logger.Debug("Closing redis monitor plugin!")
	return nil
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	p := &RedisMonitorPlugin{
		logger: logger,
	}
	var pluginMap = map[string]plugin.Plugin{
		"pilotGo_plugin": &plugin_sdk.PilotGoPlugin{Impl: p},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin_sdk.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}
type InfoResult struct {
	Server  interface{} `json:"server"`
	Comm    interface{} `json:"comm"`
	Cpu     interface{} `json:"cpu"`
	Memory  interface{} `json:"memory"`
	Stats   interface{} `json:"stats"`
	Cluster interface{} `json:"cluster"`
}

func startWebServer(g *RedisMonitorPlugin) {
	g.logger.Debug("Starting Redis Plugin server on port:25564")

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("content-type", "application/json")
		servers, comm, cpu, memory, stats, cluster, err := g.conn.get()
		if err != nil {
			return
		}
		res, _ := json.Marshal(Response{Result: InfoResult{Server: servers, Comm: comm, Cpu: cpu, Cluster: cluster, Memory: memory, Stats: stats}, Code: 200})
		fmt.Fprintf(w, string(res))
	})

	stripped, err := fs.Sub(frontend, "dist")
	if err != nil {
		g.logger.Error("start frontend fail", err.Error())
	}
	frontendFS := http.FileServer(http.FS(stripped))
	http.Handle("/", frontendFS)

	http.ListenAndServe(":25564", nil)
}
