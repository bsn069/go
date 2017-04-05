package bin

import "fmt"

var g_cmd map[string]func(strParam string, bHelp bool)
var g_lastBin IBin

func cmdReg() {
	g_cmd = make(map[string]func(strParam string, bHelp bool))
	g_cmd["LS"] = cmdLS
	g_cmd["HELP"] = cmdHELP
	g_cmd["QUIT"] = cmdQUIT
	g_cmd["CD"] = cmdCD
	g_cmd["PWD"] = cmdPWD
}

func cmdPWD(strParam string, bHelp bool) {
	if bHelp {
		fmt.Println("current bin")
		return
	}
}

func cmdQUIT(strParam string, bHelp bool) {
	if bHelp {
		fmt.Println("quit bin")
		return
	}
}

func cmdCD(strParam string, bHelp bool) {
	if bHelp {
		fmt.Println("enter bin")
		return
	}
}

func cmdLS(strParam string, bHelp bool) {
	if bHelp {
		fmt.Println("show all cmd")
		return
	}
}

func cmdHELP(strParam string, bHelp bool) {
	if bHelp {
		fmt.Println("show help")
		return
	}

	for k, v := range g_cmd {
		fmt.Print(k + "\t")
		v("", true)
	}
}
