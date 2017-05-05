package gate

func (this *SBin) AsyncCmd(strCmd, strParam string) {
	BinCmdRun(this, strCmd, strParam, false)
}
