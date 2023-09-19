package tune

type GoGcApp struct{}

func (g *GoGcApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "go_gc",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/go_gc.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf go_gc.tar.gz",
		Prepare:       "cd /tmp/tune/go_gc && sh prepare.sh",
		Tune:          "atune-adm tuning --project go_gc --detail go_gc_client.yaml",
		Restore:       "atune-adm tuning --restore --project go_gc",
	}
	return info
}
