package tune

type RedisApp struct{}

func (redis *RedisApp) Info() *TuneInfo {
	info := &TuneInfo{}
	return info
}
