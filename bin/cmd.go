package bin

import "fmt"
import "strings"
import "strconv"

func cmdReg() {
	g_cmd = make(map[string]func(strParam string, bHelp bool) bool)
	g_cmd["LS"] = cmdLS
	g_cmd["HELP"] = cmdHELP
	g_cmd["QUIT"] = cmdQUIT
	g_cmd["CD"] = cmdCD
	g_cmd["PWD"] = cmdPWD
}

func cmdPWD(strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("current bin")
		fmt.Println("\tusage: pwd")
		return true
	}

	if g_lastBin == nil {
		fmt.Println(0, 0)
	} else {
		fmt.Println(g_lastBin.TypeId(), g_lastBin.Id())
	}
	return true
}

func cmdQUIT(strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("quit bin")
		return true
	}

	quitBin()
	return true
}

func cmdCD(strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("enter bin")
		fmt.Println("\tusage: cd [(uint32)typeId (uint32)id]")
		return true
	}

	// fmt.Println(len(strParam))
	if len(strParam) == 0 {
		g_lastBin = nil
		return true
	}

	strParams := strings.Split(strParam, " ")
	if len(strParams) != 2 {
		fmt.Println("error param count")
		return false
	}

	typeId, err := strconv.ParseUint(strParams[0], 10, 32)
	if err != nil {
		fmt.Println("error typeId")
		return false
	}

	id, err := strconv.ParseUint(strParams[1], 10, 32)
	if err != nil {
		fmt.Println("error id")
		return false
	}

	iBin := getBin(uint32(typeId), uint32(id))
	if iBin == nil {
		fmt.Println("not found")
		return false
	}
	g_lastBin = iBin
	return true
}

func cmdLS(strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("show all bin")
		return true
	}
	showBin()
	return true
}

func cmdHELP(strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("show help")
		return true
	}

	// fmt.Println("in bin run system cmd use !cmd")
	for k, v := range g_cmd {
		fmt.Print(k + "\t")
		v("", true)
	}

	return true
}
