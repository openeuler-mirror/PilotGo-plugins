package tune

type FioApp struct{}

type FioImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (f *FioApp) Info() *FioImp {
	info := &FioImp{
		BaseTune: TuneInfo{
			TuneName:      "compress_Except",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/fio.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf fio.tar.gz",
			Prepare:       "cd /tmp/tune/fio && sh prepare.sh 参数一 参数二",
			Tune:          "atune-adm tuning --project compress_Except_example --detail compress_Except_example_client.yaml",
			Restore:       "atune-adm tuning --restore --project compress_Except_example",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [test_path] [test_disk]",
	}
	return info
}
