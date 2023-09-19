package tune

type IozoneApp struct{}

type IozoneImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (io *IozoneApp) Info() *IozoneImp {
	info := &IozoneImp{
		BaseTune: TuneInfo{
			TuneName:      "iozone",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/iozone.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf iozone.tar.gz",
			Prepare:       "cd /tmp/tune/fio && sh prepare.sh 参数一 参数二",
			Tune:          "atune-adm tuning --project iozone --detail tuning_iozone_client.yaml",
			Restore:       "atune-adm tuning --restore --project iozone",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [test_path] [diskname]",
	}
	return info
}
