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

	wdID uint32
}

func NewManager(dir string, id uint32) *Manager {
	fm := Manager{PWD: dir, wdID: id}
	fm.populate()
	fmt.Printf("%+v\n", fm.Files)
	return &fm
}

func (m *Manager) Icon(file *File) string {
	println(file.mimeTypeMajor, file.mimeTypeMinor)
	return strings.Split(mime[file.mimeTypeMajor][file.mimeTypeMinor], "/")[1]
}

func (m *Manager) populate() {
	// get info from two sources:
	// 1. every content entity that has f.wd as parent
	// 2. inspect the actual os level file system
	m.Files = fromStorageChildren(m.wdID)
}
