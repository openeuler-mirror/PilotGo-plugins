package tune

type GccComplieApp struct{}

func (gcc *GccComplieApp) Info() *TuneInfo {
	info := &TuneInfo{
		Prepare: "mkdir -p /tmp/tune/ && cp -r ../../templete/gcc_compile /tmp/tune/ && sh prepare.sh",
		Tune:    "atune-adm tuning --project gcc_compile --detail gcc_compile_client.yaml",
		Restore: "atune-adm tuning --restore --project gcc_compile",
	}
	return info
}
func (gcc *GccComplieApp) Prepare() error {
	return nil
}
func (gcc *GccComplieApp) StartTune() error {
	return nil
}
func (gcc *GccComplieApp) Restore() error {
	return nil
}
