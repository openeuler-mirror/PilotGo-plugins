package tune

type MariadbApp struct{}

type MariadbImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (m *MariadbApp) Info() *MariadbImp {
	info := &MariadbImp{
		BaseTune: TuneInfo{
			TuneName:      "mariadb",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/mariadb.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf mariadb.tar.gz",
			Prepare:       "cd /tmp/tune/mariadb && sh prepare.sh 25",
			Tune:          "atune-adm tuning --project mariadb --detail mariadb_client.yaml",
			Restore:       "atune-adm tuning --restore --project mariadb",
		},
		Notes: "请注意prepare.sh的使用方法: sh prepare.sh [迭代次数, 默认25]",
	}
	return info
}
