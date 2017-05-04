package bin

import (
	"time"
)

type IBin interface {
	AsyncCmd(strCmd, strParam string)
	Update(u32ElapseMs uint32, now *time.Time) bool
	TypeId() uint32
	Id() uint32
	Quit()
	FrameMs() uint32 // every frame max ms
}

func Run(iBin IBin) {
	run(iBin)
}

func WaitQuit() {
	waitQuit()
}
