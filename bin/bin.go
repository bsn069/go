package bin

func run(iBin IBin) {

	if iBin == nil {
		return
	}

	if addBin(iBin) != nil {
		return
	}

	for {
		if !iBin.ProcCmd("") {
			break
		}

		if !iBin.Update() {
			break
		}
	}
}
