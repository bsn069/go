package gate

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bsn069/go/bin"
)

func BinCmdRun(pBin *SBin, strCmd, strParam string, bHelp bool) {
	strCmd = strings.ToUpper(strCmd)

	if len(strCmd) == 0 { // show help
		binCmdHelp(pBin, strParam, false)
		return
	}

	if strCmd[0] == '?' { // show cmd help
		BinCmdRun(pBin, strCmd[1:], strParam, true)
		return
	}

	if funcExec, bOk := g_binCmd[strCmd]; bOk {
		if !funcExec(pBin, strParam, bHelp) && !bHelp {
			funcExec(pBin, strParam, true)
		}
		return
	}

	// find cmd from base bin
	bOk := bin.BinCmdRun(&pBin.SBinBase, strCmd, strParam, bHelp)
	if !bOk {
		// not found any cmd, run help
		binCmdHelp(pBin, strParam, false)
	}
}

func binCmdReg(strCmd string, funcCmd func(pBin *SBin, strParam string, bHelp bool) bool) {
	strCmd = strings.ToUpper(strCmd)
	g_binCmd[strCmd] = funcCmd
}

func binCmdHelp(pBin *SBin, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\tshow help")
		return true
	}

	fmt.Println("=======================")
	fmt.Println("run system cmd use !cmd|")
	fmt.Println("show cmd help use ?cmd |")
	fmt.Println("=======================")

	// fmt.Println("Help for SBin")
	var buffer bytes.Buffer
	for k, v := range g_binCmd {
		buffer.WriteString(k)
		buffer.WriteString(" ")
		fmt.Println(k)
		v(nil, "", true)
	}
	strParam = buffer.String()
	bin.BinCmdRun(&pBin.SBinBase, "HELP", strParam, bHelp)
	return true
}

func binCmdTest(pBin *SBin, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\ttest cmd")
		return true
	}

	return true
}
