package widget

const menID = "id"
const menName = "name"
const menEntries = "entry_ids"

const menuKeys = menID + "," + menName + "," + menEntries

type Menu struct {
	ID      uint16
	Name    string
	Entries []*MenuEntry
}

type MenuEntry struct {
	ID       string
	PageName string
}

func (m *Menu) ScanFromDB(attr string, val interface{}) bool {
	return false
}

func (m *Menu) Markup() string {
	var mk string

	for _, entry := range m.Entries {
		mk += "<li>" + entry.PageName + "</li>"
	}

	//return "<ul>" + mk + "</ul>"
	return mk
}

func (m *Menu) Store() {

}
