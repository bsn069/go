package bin

import (
	"fmt"
	"time"
)

func run(iBin IBin) {
	g_waitQuit.Add(1)
	go runThread(iBin)
}

func runThread(iBin IBin) {
	defer g_waitQuit.Done()

	if iBin == nil {
		return
	}

	if addBin(iBin) != nil {
		return
	}

	for {
		time.Sleep(time.Duration(time.Second * 1))
		if !iBin.Update() {
			break
		}
	}
	fmt.Println("had quit")
}

func waitQuit() {
	g_waitQuit.Wait()
}
