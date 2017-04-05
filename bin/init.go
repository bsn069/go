package bin

import (
	"errors"
	"sync"
)

var g_bins map[uint64]IBin
var g_binMutex sync.Mutex

func init() {
	cmdReg()
	g_bins = make(map[uint64]IBin)
	go inputThread()
}

func makeBinKey(iBin IBin) uint64 {
	typeId := iBin.TypeId()
	id := iBin.Id()
	key := uint64(typeId)
	key <<= 32
	key |= uint64(id)
	return key
}

func addBin(iBin IBin) error {
	key := makeBinKey(iBin)

	g_binMutex.Lock()
	defer g_binMutex.Unlock()

	_, ok := g_bins[key]
	if ok {
		return errors.New("had add key ")
	}
	g_bins[key] = iBin
	return nil
}
