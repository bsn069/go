package bin

type SBinTest struct {
	SBinBase
}

func NewBinTest(typeid, id uint32) *SBinTest {
	pBin := &SBinTest{}
	InitBinBase(typeid, id, &pBin.SBinBase)
	pBin.SetFrameMs(11111111)
	return pBin
}
