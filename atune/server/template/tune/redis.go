package tune

type RedisApp struct{}

type RedisImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (redis *RedisApp) Info() *RedisImp {
	info := &RedisImp{
		BaseTune: TuneInfo{
			TuneName:      "redis",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/redis.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf redis.tar.gz",
			Prepare:       "cd /tmp/tune/redis && sh prepare.sh",
			Tune:          "atune-adm tuning --project redis --detail ./redis_client.yaml",
			Restore:       "atune-adm tuning --restore --project redis",
		},
		Notes: "若要启动基准测试, 在调优前先进入调优工作目录/tmp/tune/redis执行“sh redis_benchmark.sh”,本地主机将访问基准测试主机并触发它。基准测试主机将在基准测试之后将日志文件传输到localhost。",
	}
	return info
}
