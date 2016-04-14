package file

const (
	Directory = 0
	TextHtml  = 1
	TextPlain = 2
	ImageJpeg = 3
	ImagePng  = 4
)

var mime = [5]string{
	"directory/folder",
	"text/html",
	"text/plain",
	"image/jpeg",
	"image/png"}
