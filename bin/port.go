package bin

type IBin interface {
	ProcCmd(string) bool
	Update() bool
	TypeId() uint32
	Id() uint32
}

func Run(iBin IBin) {
	run(iBin)
}

type SBinTest struct {
}

func (this *SBinTest) ProcCmd(string) bool {
	return true
}
func (this *SBinTest) Update() bool {
	return true
}
func (this *SBinTest) TypeId() uint32 {
	return 1
}
func (this *SBinTest) Id() uint32 {
	return 2
}
