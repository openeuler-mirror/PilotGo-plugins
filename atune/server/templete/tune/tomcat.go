package tune

type TomcatApp struct{}

func (tomcat *TomcatApp) Info() *TuneInfo {
	info := &TuneInfo{}
	return info
}
