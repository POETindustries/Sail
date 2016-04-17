package file

import "fmt"

const (
	Private = 0
	Public  = 1
)

type File struct {
	ID      uint32
	Name    string
	Address string

	machineName   string
	parent        uint32
	owner         uint32
	mimeTypeMajor uint16
	mimeTypeMinor uint16
	status        int8
	cDate         string
	eDate         string
}

func (f *File) Status() string {
	if f.status == 0 {
		return "private"
	}
	return "public"
}

func (f *File) String() string {
	return fmt.Sprintf("FILE %s: {Address: %s | Type: %s | Status: %s}",
		f.Name, f.Address, mime[f.mimeTypeMajor][f.mimeTypeMinor], f.Status())
}

func (f *File) Type() string {
	return mime[f.mimeTypeMajor][f.mimeTypeMinor]
}

func (f *File) TypeCode() (uint16, uint16) {
	return f.mimeTypeMajor, f.mimeTypeMinor
}
