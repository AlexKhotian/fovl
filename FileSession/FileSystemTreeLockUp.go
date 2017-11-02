package FileSession

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileSystemTreeLookUp is used to search for files in folder
type FileSystemTreeLookUp struct {
}

// SearchForFileItemsInDirectory used to find files in dir
func (lookUpHandler *FileSystemTreeLookUp) SearchForFileItemsInDirectory(dir string) []string {
	var fileItemsInDir []string

	var walkFunc filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return err
	}
	filepath.Walk(dir, walkFunc)
	return fileItemsInDir
}
