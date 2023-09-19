package tune

type NginxApp struct{}

type NginxImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (nginx *NginxApp) Info() *NginxImp {
	info := &NginxImp{
		BaseTune: TuneInfo{
			TuneName:      "nginx",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/nginx.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf nginx.tar.gz",
			Prepare:       "cd /tmp/tune/nginx && sh prepare.sh",
			Tune:          "atune-adm tuning --project nginx --detail nginx_client.yaml",
			Restore:       "atune-adm tuning --restore --project nginx",
		},
		Notes: "nginx长连接:【调优指令】atune-adm tuning --project nginx_http_long --detail nginx_http_long_client.yaml【恢复环境】atune-adm tuning --restore --project nginx_http_long",
	}
	return info
}
