package tune

type MemcachedApp struct{}

func (m *MemcachedApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "memcached",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/memcached.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf memcached.tar.gz",
		Prepare:       "cd /tmp/tune/memcached && sh prepare.sh",
		Tune:          "atune-adm tuning --project memcached_memaslap --detail memcached_memaslap_client.yaml",
		Restore:       "atune-adm tuning --restore --project memcached_memaslap",
	}
	return info
}
