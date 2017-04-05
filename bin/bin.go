package bin

func run(iBin IBin) {

	if iBin == nil {
		return
	}

	if addBin(iBin) != nil {
		return
	}

	for {
		if !iBin.Update() {
			break
		}
	}
}
