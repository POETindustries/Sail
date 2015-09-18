package page

import (
	"sail/core/dbase"
)

const domID = "id"
const domName = "name"
const domTemplate = "template"

const domainKeys = domID + "," + domName + "," + domTemplate

type Domain struct {
	ID       uint32
	Name     string
	Template string
}

func (d *Domain) ScanFromDB(conn *dbase.Conn) bool {
	data := conn.DomainData(domainKeys, d.ID)

	if data.Scan(&d.ID, &d.Name, &d.Template) != nil {
		return false
	}

	return true
}
