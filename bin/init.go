package bin

import (
	"sync"
)

var g_bins map[uint64]IBin
var g_binMutex sync.Mutex
var g_waitQuit sync.WaitGroup
var g_cmd map[string]func(strParam string, bHelp bool) bool
var g_lastBin IBin

func init() {
	cmdReg()
	g_bins = make(map[uint64]IBin)
	go inputThread()
}
