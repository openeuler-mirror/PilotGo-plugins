package tune

type MemoryApp struct{}

func (m *MemoryApp) Info() *TuneInfo {
	info := &TuneInfo{}
	return info
}
