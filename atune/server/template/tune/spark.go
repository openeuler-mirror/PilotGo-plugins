package tune

type SparkApp struct{}

type SparkImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (spark *SparkApp) Info() *SparkImp {
	info := &SparkImp{
		BaseTune: TuneInfo{
			TuneName:      "spark",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/spark.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf spark.tar.gz",
			Prepare:       "cd /tmp/tune/spark && sh run_env.sh",
			Tune:          "",
			Restore:       "",
		},
		Notes: "调优步骤均在run_env.sh完成, 执行该脚本可完成环境的部署和调优测试",
	}
	return info
}
