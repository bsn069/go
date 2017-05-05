package gate

var g_binCmd map[string]func(pBin *SBin, strParam string, bHelp bool) bool

func baseInit() {
	g_binCmd = make(map[string]func(pBin *SBin, strParam string, bHelp bool) bool)
	binCmdReg("help", binCmdHelp)
	binCmdReg("test", binCmdTest)
}
