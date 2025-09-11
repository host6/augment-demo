package filesu

import "io/fs"

const (
	FileMode_rwxrwxrwx fs.FileMode = 0777 // default for directory
	FileMode_rw_rw_rw_ fs.FileMode = 0666 // default for file
)
