// Package file handles the presentation of content and
// other objects to the user.
//
// It provides a generic File object that exposes object's
// meta data and more generic attributes shared across
// objects of different types.
//
// File objects are usually used in conjunction with an
// instance of file.Manager. The Manager object organizes
// files in a familiar way, letting users traverse through
// directories and select files for editing, deletion or
// any other CRUD operation.
//
// Note that the term 'file' is used in UNIX way, meaning
// that even directories are represented by the File type.
// What differentiates folders from other files is the
// state of relevant attributes.
package file
