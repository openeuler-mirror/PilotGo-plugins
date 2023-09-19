package tune

type GraphicsmagickApp struct{}

func (g *GraphicsmagickApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "graphicsmagick",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/graphicsmagick.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf graphicsmagick.tar.gz",
		Prepare:       "cd /tmp/tune/graphicsmagick && sh prepare.sh",
		Tune:          "atune-adm tuning --project graphicsmagick  --detail gm_client.yaml",
		Restore:       "atune-adm tuning --restore --project graphicsmagick",
	}
	return info
}
