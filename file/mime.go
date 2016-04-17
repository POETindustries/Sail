package file

const (
	Directory = 0
	Folder    = 0
)

const (
	Text  = 1
	Html  = 0
	Plain = 1
)

const (
	Image = 2
	Jpeg  = 0
	Png   = 1
)

var mime = [3][]string{
	[]string{"directory/folder"},
	[]string{"text/html", "text/plain"},
	[]string{"image/jpeg", "image/png"}}
