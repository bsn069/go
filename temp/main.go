package main

import (
	_ "net/http/pprof"

	"github.com/bsn069/go/bin"
)

func main() {
	binTest := new(bin.SBinTest)
	bin.Run(binTest)
}
