package tune

type MysqlSysbenchApp struct{}

type MysqlSysbenchImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (m *MysqlSysbenchApp) Info() *MysqlSysbenchImp {
	info := &MysqlSysbenchImp{
		BaseTune: TuneInfo{
			TuneName:      "mysql_sysbench",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/mysql_sysbench.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf mysql_sysbench.tar.gz",
			Prepare:       "cd /tmp/tune/mysql_sysbench && sh prepare.sh",
			Tune:          "atune-adm tuning --project mysql_sysbench --detail mysql_sysbench_client.yaml",
			Restore:       "atune-adm tuning --restore --project mysql_sysbench",
		},
		Notes: "1,根据指南安装mysql: https://blog.csdn.net/weixin_43214408/article/details/116895091 \n 2,根据指南安装sysbench: https://blog.csdn.net/weixin_43214408/article/details/116898751",
	}
	return info
}
