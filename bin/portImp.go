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

	var timeElapseMs uint32
	for {
		timeNow := time.Now()
		timeBegin := timeNow.UnixNano()
		if !iBin.Update(timeElapseMs, &timeNow) {
			break
		}
		timeEnd := time.Now().UnixNano()

		frameMs := iBin.FrameMs()
		timeElapse := timeEnd - timeBegin
		if timeElapse < 0 { // time back
			timeElapseMs = frameMs
		} else {
			timeElapseMs = uint32(timeElapse / int64(time.Millisecond))
		}

		if timeElapseMs < frameMs {
			timeWaitMs := frameMs - timeElapseMs
			time.Sleep(time.Duration(time.Millisecond * time.Duration(timeWaitMs)))
			timeElapseMs += timeWaitMs
		}
	}
	fmt.Println("had quit")
}

func waitQuit() {
	g_waitQuit.Wait()
}
