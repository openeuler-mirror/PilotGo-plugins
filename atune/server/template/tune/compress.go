package tune

type CompressApp struct{}

func (c *CompressApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "compress",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/compress.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf compress.tar.gz",
		Prepare:       "cd /tmp/tune/compress && sh prepare.sh enwik8.zip",
		Tune:          "atune-adm tuning --project compress --detail compress_client.yaml",
		Restore:       "atune-adm tuning --restore --project compress",
	}
	return info
}
