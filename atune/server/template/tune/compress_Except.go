package tune

type CompressExceptApp struct{}

func (ce *CompressExceptApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "compress_Except",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/compress_Except_example.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf compress_Except_example.tar.gz",
		Prepare:       "cd /tmp/tune/compress_Except_example && sh prepare.sh enwik8.zip",
		Tune:          "atune-adm tuning --project compress_Except_example --detail compress_Except_example_client.yaml",
		Restore:       "atune-adm tuning --restore --project compress_Except_example",
	}
	return info

}
