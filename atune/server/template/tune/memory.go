package tune

type MemoryApp struct{}

func (m *MemoryApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "memory",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/memory.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf memory.tar.gz",
		Prepare:       "cd /tmp/tune/memory && sh prepare.sh",
		Tune:          "atune-adm tuning --project stream --detail tuning_stream_client.yaml",
		Restore:       "atune-adm tuning --restore --project stream",
	}
	return info
}
