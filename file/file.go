package file

import "fmt"

// Private and Public are the numerical representations for
// a file's public visibility status.
const (
	Private = 0
	Public  = 1
)

// File represents an object (content object, folder, static
// image file...) in a way that the file manager understands
// and can work with.
//
// Note that the term 'file' is used in a UNIX way, i.e.
// directories are files, too.
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

// IsDir returns true if the file is a directory.
func (f *File) IsDir() bool {
	return f.mimeTypeMajor == Directory
}

// Status returns a human-readable representation of the
// file's pubic visibility.
func (f *File) Status() string {
	if f.status == 0 {
		return "private"
	}
	return "public"
}

// StatusCode returns a machine-readable representation of
// the file's public visibility.
func (f *File) StatusCode() int8 {
	return f.status
}

// String returns a formatted string containing information
// about the file's current state.
func (f *File) String() string {
	return fmt.Sprintf("FILE %s: {Address: %s | Type: %s | Status: %s}",
		f.Name, f.Address, mime[f.mimeTypeMajor][f.mimeTypeMinor], f.Status())
}

// Type returns a mime type expression representing the file's
// mime type.
func (f *File) Type() string {
	return mime[f.mimeTypeMajor][f.mimeTypeMinor]
}

// TypeCode returns the internally assigned major and minor
// mime type of the file.
func (f *File) TypeCode() (uint16, uint16) {
	return f.mimeTypeMajor, f.mimeTypeMinor
}

// TypeCodeMajor returns only the major part of the file's
// mime type, represented by an internally assigned number.
func (f *File) TypeCodeMajor() uint16 {
	return f.mimeTypeMajor
}

// TypeCodeMinor returns only the minor part of the file's
// mime type, represented by an internally assigned number.
func (f *File) TypeCodeMinor() uint16 {
	return f.mimeTypeMinor
}

// hasChildren looks into the database and returns true if
// any object has its parent point to this file.
func (f *File) hasChildren() bool {
	return fromStorageChildCount(f.ID) > 0
}
