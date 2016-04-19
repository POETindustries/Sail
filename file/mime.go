package file

// Directory is the major mime type for directories.
const (
	Directory = 0
	Folder    = 0
)

// Text is the major mime type for human-readable text files.
const (
	Text  = 1
	Html  = 0
	Plain = 1
)

// Image is the major mime type for image files.
const (
	Image = 2
	Jpeg  = 0
	Png   = 1
)

var mime = [3][]string{
	[]string{"directory/folder"},
	[]string{"text/html", "text/plain"},
	[]string{"image/jpeg", "image/png"}}
