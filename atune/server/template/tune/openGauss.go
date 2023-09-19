package tune

type OpenGaussApp struct{}

type OpenGaussImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (gauss *OpenGaussApp) Info() *OpenGaussImp {
	info := &OpenGaussImp{
		BaseTune: TuneInfo{
			TuneName:      "openGauss",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/openGauss.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf openGauss.tar.gz",
			Prepare:       "cd /tmp/tune/openGauss && sh prepare.sh",
			Tune:          "atune-adm tuning --project openGauss_tpcc --detail openGauss_client.yaml",
			Restore:       "atune-adm tuning --restore --project openGauss_tpcc",
		},
		Notes: "请先安装openGauss和benchmarksql-5.0 \n openGauss安装指南: https://docs.opengauss.org/zh/docs/2.1.0/docs/installation/installation.html \n benchmarksql-5.0下载地址: https://udomain.dl.sourceforge.net/project/benchmarksql/benchmarksql-5.0.zip",
	}
	return info
}
