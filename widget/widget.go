package widget

import (
	"sail/widget/dbase"
)

type Widget interface {
	ScanFromDB(attr string, val interface{}) bool
}

func Init() {
	dbase.Init()
}
