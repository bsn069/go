package bin

import (
	"errors"
	"fmt"
)

func getBinKey(iBin IBin) uint64 {
	typeId := iBin.TypeId()
	id := iBin.Id()
	return makeBinKey(typeId, id)
}

func makeBinKey(typeId, id uint32) uint64 {
	key := uint64(typeId)
	key <<= 32
	key |= uint64(id)
	return key
}

func addBin(iBin IBin) error {
	key := getBinKey(iBin)

	g_binMutex.Lock()
	defer g_binMutex.Unlock()

	_, ok := g_bins[key]
	if ok {
		return errors.New("had add key ")
	}
	g_bins[key] = iBin
	return nil
}

func showBin() {
	for _, v := range g_bins {
		fmt.Println(v.TypeId(), v.Id())
	}
}

func quitBin() {
	for _, v := range g_bins {
		v.Quit()
	}
}

func getBin(typeId, id uint32) IBin {
	key := makeBinKey(typeId, id)
	iBin, ok := g_bins[key]
	if !ok {
		return nil
	}
	return iBin
}
