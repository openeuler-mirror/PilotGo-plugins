package pluginclient

import (
	"sync"

	"openeuler.org/PilotGo/plugin-sdk/plugin"
)

var once sync.Once
var global_client *plugin.Client

func NewClient(desc *plugin.PluginInfo) {
	once.Do(func() {
		// client init
		global_client = plugin.DefaultClient(desc)
	})
}

func Client() *plugin.Client {
	return global_client
}
