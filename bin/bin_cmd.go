package bin

import (
	"fmt"
	"strconv"
	"strings"
)

func BinCmdRun(pBin *SBinBase, strCmd, strParam string, bHelp bool) bool {
	strCmd = strings.ToUpper(strCmd)
	if funcExec, bOk := g_binCmd[strCmd]; bOk {
		if !funcExec(pBin, strParam, bHelp) && !bHelp {
			funcExec(pBin, strParam, true)
		}
		return true
	}
	return false
}

func binCmdReg() {
	g_binCmd = make(map[string]func(pBin *SBinBase, strParam string, bHelp bool) bool)
	g_binCmd["SETFRAMEMS"] = binCmdSetFrameMs
	g_binCmd["SHOWFRAMEMS"] = binCmdShowFrameMs
	g_binCmd["HELP"] = binCmdHelp
}

func binCmdSetFrameMs(pBin *SBinBase, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\tset frame time ms")
		return true
	}

	strParams := strings.Split(strParam, " ")
	if len(strParams) != 1 {
		fmt.Println("error param count")
		return false
	}

	uMs, err := strconv.ParseUint(strParams[0], 10, 32)
	if err != nil {
		fmt.Println("error ms")
		return false
	}
	pBin.SetFrameMs(uint32(uMs))
	return true
}

func binCmdShowFrameMs(pBin *SBinBase, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\tshow frame time ms")
		return true
	}

	fmt.Println("frameMs =", pBin.M_u32FrameMs)
	return true
}

func binCmdHelp(pBin *SBinBase, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\tshow help")
		return true
	}

	noShowCmd := make(map[string]struct{})
	strParams := strings.Split(strParam, " ")
	for _, v := range strParams {
		noShowCmd[v] = struct{}{}
	}

	// fmt.Println("Help for SBinBase")
	for k, v := range g_binCmd {
		if _, bOk := noShowCmd[k]; bOk {
			continue
		}
		fmt.Println(k)
		v(nil, "", true)
	}

	return true
}
