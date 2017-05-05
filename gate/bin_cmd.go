package gate

import (
	"fmt"
)

func binCmdTest1(pBin *SBin, strParam string, bHelp bool) bool {
	if bHelp {
		fmt.Println("\ttest1 cmd")
		return true
	}

	return true
}
