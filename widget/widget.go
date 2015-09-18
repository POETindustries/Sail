package widget

type Widget interface {
	ScanFromDB(attr string, val interface{}) bool
}
