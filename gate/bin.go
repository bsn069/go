package gate

import (
	"github.com/bsn069/go/bin"
)

type SBin struct {
	bin.SBinBase
}

func NewBin(id uint32) *SBin {
	pBin := &SBin{}
	bin.InitBinBase(99, id, &pBin.SBinBase)
	return pBin
}
