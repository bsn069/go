package bin

import "fmt"

var g_cmd map[string]func(strParam string, bHelp bool)

func cmdReg() {
	g_cmd = make(map[string]func(strParam string, bHelp bool))
	g_cmd["LS"] = cmdLS
	g_cmd["HELP"] = cmdHELP
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
