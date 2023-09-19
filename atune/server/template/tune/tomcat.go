package tune

type TomcatApp struct{}

type TomcatImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (tomcat *TomcatApp) Info() *TomcatImp {
	info := &TomcatImp{
		BaseTune: TuneInfo{
			TuneName:      "tomcat",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/tomcat.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf tomcat.tar.gz",
			Prepare:       "cd /tmp/tune/tomcat && sh prepare.sh tomcat_root",
			Tune:          "atune-adm tuning --project tomcat --detail tomcat.yaml",
			Restore:       "atune-adm tuning --restore --project tomcat",
		},
		Notes: "注意prepare.sh的使用方法, sh prepare.sh [tomcat 目录名称], 默认是tomcat_root",
	}
	return info
}
