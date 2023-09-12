package tune

type NginxApp struct{}

func (nginx *NginxApp) Info() *TuneInfo {
	info := &TuneInfo{}
	return info
}
