/*
func Test_Input(t *testing.T) {
	count := 3
	for index := 0; index < count; index++ {
		time.Sleep(time.Duration(5 * time.Second))
		strLines := ReadLines()
		for strLine := range strLines {
			fmt.Println(strLine)
		}
	}
}
*/
package bin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func inputThread() {
	r := bufio.NewReader(os.Stdin)
	for {
		byLine, _, _ := r.ReadLine()
		strLine := string(byLine)
		// fmt.Println(strLine)
		var strCmd string = strLine
		var strParam string
		for k, v := range strLine {
			// fmt.Println(k, v)
			if unicode.IsSpace(v) {
				strCmd = strLine[0:k]
				strParam = strLine[k+1:]
				break
			}
		}

		if len(strCmd) >= 1 {
			if strCmd[0] == '!' {
				cmdExec(strCmd[1:], strParam)
			} else {

			}
		}
	}
}

func cmdExec(strCmd, strParam string) {
	strCmd = strings.ToUpper(strCmd)
	if funcExec, bOk := g_cmd[strCmd]; bOk {
		fmt.Println("run " + strCmd)
		funcExec(strParam, false)
	} else {
		cmdHELP(strParam, false)
	}
}
