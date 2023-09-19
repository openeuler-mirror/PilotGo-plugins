package tune

type TidbApp struct{}

type TidbImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (tidb *TidbApp) Info() *TidbImp {
	info := &TidbImp{
		BaseTune: TuneInfo{
			TuneName:      "tidb",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/tidb.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf tidb.tar.gz",
			Prepare:       "cd /tmp/tune/tidb && sh prepare.sh",
			Tune:          "atune-adm tuning --project tidb --detail tidb_client.yaml",
			Restore:       "atune-adm tuning --restore --project tidb",
		},
		Notes: "如果目标机器是第一次使用tidb调优, 请为目标机器准备好测试环境, 详情请查看 https://gitee.com/openeuler/A-Tune/blob/master/examples/tuning/tidb/README",
	}
	return info
}
