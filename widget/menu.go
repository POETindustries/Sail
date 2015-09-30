package widget

import "html/template"

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
	ID       uint32
	PageName string
}

func (m *Menu) ScanFromDB(attrs string, val interface{}) bool {
	/*	data := storage.MenuData(attrs, val)
		var pageIDs string

		if err := data.Scan(&m.ID, &m.Name, &pageIDs); err != nil {
			println("Here?")
			plugin.LogError(err, plugin.ConfigData().DevMode)
			println("Here.")
			return false
		}

		m.deserialize(pageIDs)*/

	return true
}

func (m *Menu) Markup() template.HTML {
	var mk string

	for _, entry := range m.Entries {
		mk += "<li>" + entry.PageName + "</li>"
	}

	//return template.HTML("<ul>" + mk + "</ul>")
	return template.HTML(mk)
}

func (m *Menu) Store() {

}

func (m *Menu) deserialize(vals string) {
	/*	slice := strings.Split(vals, ",")
		m.Entries = make([]*MenuEntry, len(slice))
		for _, val := range slice {
			id, _ := strconv.Atoi(val)
			entry := MenuEntry{ID: uint32(id)}
			data := storage.PageName(uint32(id))
			if err := data.Scan(&entry.PageName); err != nil {
				plugin.LogError(err, plugin.ConfigData().DevMode)
			}
			m.Entries = append(m.Entries, &entry)
		}*/
}
