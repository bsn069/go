package bin

type IBin interface {
	AsyncCmd(strCmd, strParam string)
	Update() bool
	TypeId() uint32
	Id() uint32
	Quit()
}

func Run(iBin IBin) {
	run(iBin)
}

func WaitQuit() {
	waitQuit()
}
