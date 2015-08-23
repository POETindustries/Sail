package errors

func Log(e error, devMode bool) {
	if devMode {
		println(e.Error())
	} else {
		// TODO more intricate error logging
	}
}
