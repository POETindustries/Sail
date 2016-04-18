package file

import (
	"net/url"
	"sail/object"
	"sail/object/cache"
	"strconv"
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

func NewManager(query url.Values) (fm *Manager) {
	id, err := strconv.ParseUint(query.Get("loc"), 10, 32)
	if err != nil {
		fm = &Manager{PWD: "/", wdID: 0}
	} else {
		uuid := "uuid/" + query.Get("loc")
		pwd := cache.DB().ObjectURL(uuid)
		if pwd == "" {
			pwd = object.StaticAddr(uuid)
		}
		fm = &Manager{PWD: pwd, wdID: uint32(id)}
	}
	fm.populate()
	return fm
}

func (m *Manager) Icon(file *File) string {
	return strings.Split(mime[file.mimeTypeMajor][file.mimeTypeMinor], "/")[1]
}

func (m *Manager) populate() {
	files := fromStorageChildren(m.wdID, true)
	var ds []*File
	var fs []*File
	for _, f := range files {
		if f.IsDir() || f.hasChildren() {
			if f.ID == m.wdID {
				f.Name = "Index"
				fs = append([]*File{f}, fs...)
				continue
			} else if !f.IsDir() {
				f.mimeTypeMajor = Directory
				f.mimeTypeMinor = Folder
			}
			ds = append(ds, f)
		} else {
			fs = append(fs, f)
		}
	}
	m.Files = append(ds, fs...)
}
