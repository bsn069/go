package bin

import (
	"sync"
)

var g_bins map[uint64]IBin
var g_binMutex sync.Mutex
var g_waitQuit sync.WaitGroup
var g_cmd map[string]func(strParam string, bHelp bool) bool
var g_lastBin IBin
var g_binCmd map[string]func(pBin *SBinBase, strParam string, bHelp bool) bool

func init() {
	cmdReg()
	binCmdReg()
	g_bins = make(map[uint64]IBin)
	go inputThread()
}
