package main

import (
	_ "net/http/pprof"

	"github.com/bsn069/go/bin"
)

func main() {
	bin.Run(bin.NewBinTest(1, 1))
	bin.Run(bin.NewBinTest(2, 1))
	bin.Run(bin.NewBinTest(2, 2))
	bin.Run(bin.NewBinTest(3, 1))
	bin.Run(bin.NewBinTest(3, 2))
	bin.Run(bin.NewBinTest(3, 3))
	bin.Run(bin.NewBinTest(9, 9))
	bin.WaitQuit()
}
