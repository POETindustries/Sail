package file

import (
	"io/ioutil"
	"net/url"
	"sail/conf"
	"sail/errors"
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
	// get info from two sources:
	// 1. every content entity that has f.wd as parent
	// 2. inspect the actual os level file system
	m.populateWithContent()
	m.populateWithStaticFiles()
}

func (m *Manager) populateWithContent() {
	m.Files = fromStorageChildren(m.wdID, true)
	first := 0
	for i, f := range m.Files {
		if f.hasChildren() {
			if f.ID == m.wdID {
				f.Name = "Index"
				first = i
			} else {
				f.mimeTypeMajor = Directory
				f.mimeTypeMinor = Folder
			}
		}
	}
	if first != 0 {
		old := make([]*File, len(m.Files))
		copy(old, m.Files)
		old = append(old[:first], old[first+1:]...)
		m.Files = append([]*File{m.Files[first]}, old...)
	}
}

func (m *Manager) populateWithStaticFiles() {
	files, err := ioutil.ReadDir(conf.Instance().FileDir + m.PWD[1:])
	if err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return
	}
	var ds, fs []*File
	for _, f := range files {
		// TODO:	if f is directory, append to ds, else determine
		// 			mime type and append to fs.
		file := File{Name: f.Name(), Address: m.PWD + f.Name()}
		if f.IsDir() {
			ds = append(ds, &file)
		} else {
			file.mimeTypeMajor = Text
			file.mimeTypeMinor = Html
			fs = append(fs, &file)
		}
	}
	m.Files = append(m.Files, ds...)
	m.Files = append(m.Files, fs...)
}
