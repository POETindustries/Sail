package file

import (
	"fmt"
	"strings"
)

// Manager represents website's content as a collection
// of browsable directories and files. Its purpose is the
// display of content in a form that makes it easy to add,
// edit and otherwise manage content in a familiar manner.
type Manager struct {
	PWD   string
	Files []*File
}

func NewManager(dir string) *Manager {
	fm := Manager{PWD: dir}
	fm.populate()
	fmt.Printf("%+v\n", fm.Files)
	return &fm
}

func (f *Manager) Icon(mimeType uint16) string {
	return strings.Split(mime[mimeType], "/")[1]
}

func (f *Manager) populate() {
	// get info from two sources:
	// 1. every content entity that has f.wd as parent
	// 2. inspect the actual os level file system
	f.Files = WebPages(f.PWD)
}
