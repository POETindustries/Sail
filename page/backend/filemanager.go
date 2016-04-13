package backend

// FileManager represents website's content as a collection
// of browsable directories and files. Its purpose is the
// display of content in a form that makes it easy to add,
// edit and otherwise manage content in a familiar manner.
type FileManager struct {
	wd string
}

// PWD returns the current directory name.
func (f *FileManager) PWD() string {
	return f.wd
}
