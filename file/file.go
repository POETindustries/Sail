package file

import "fmt"

const (
	Private = 0
	Public  = 1
)

type File struct {
	Name    string
	Address string

	mimeType uint16
	status   int8
}

func (f *File) Status() string {
	if f.status == 0 {
		return "private"
	}
	return "public"
}

func (f *File) String() string {
	return fmt.Sprintf("FILE %s: {Address: %s | Type: %s | Status: %s}",
		f.Name, f.Address, mime[f.mimeType], f.Status())
}

func (f *File) Type() string {
	return mime[f.mimeType]
}

func (f *File) TypeCode() uint16 {
	return f.mimeType
}

func StaticAddr(uuid string) string {
	if a := fromStorageGetAddr(uuid, true); a != "" {
		return a
	}
	return uuid
}

func WebPages(dir string) []*File {
	return fromStorageAsContent(dir)
}
