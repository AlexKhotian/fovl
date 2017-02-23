package FileSession

import (
	"container/list"
	"log"
	"os"
	"strings"
)

// IFilePathHelper stores a path, behaves as normal filepath
type IFilePathHelper interface {
	// checks
	IsDirectory() bool
	IsAbsolutePath() bool
	IsRelativePath() bool
	IsFilePathExists() bool

	// Getters
	GetRelativePath() string
	GetAboslutePath() string
	GetParentName() string
	GetFileName() string
	GetParentAbsolutePath() string
	GetChildItems() *list.List

	// mutators for existing string
	AppendPathToExisting(additionalPath *string)
}

// filePathHelper impls
type filePathHelper struct {
	_filePathString *string
}

//
func CreateFilePathHelper(filePath *string) IFilePathHelper {
	this := new(filePathHelper)
	this._filePathString = filePath
	return this
}

func (helper *filePathHelper) IsDirectory() bool {
	file, err := os.Stat(*helper._filePathString)
	if err != nil {
		log.Fatal("Failed to create file")
	}
	return file.IsDir()
}

func (helper *filePathHelper) IsAbsolutePath() bool {
	return strings.HasPrefix(*helper._filePathString, "/")
}

func (helper *filePathHelper) IsRelativePath() bool {
	return !helper.IsAbsolutePath()
}

func (helper *filePathHelper) IsFilePathExists() bool {
	_, err := os.Stat(*helper._filePathString)
	if err != nil {
		return false
	}
	return true
}

func (helper *filePathHelper) GetRelativePath() string {
	slittedPath := strings.Split(*helper._filePathString, "/")
	return slittedPath[len(slittedPath)-1]
}

func (helper *filePathHelper) GetAboslutePath() string {
	return *helper._filePathString
}

func (helper *filePathHelper) GetParentName() string {
	slittedPath := strings.Split(*helper._filePathString, "/")
	return slittedPath[len(slittedPath)-2]
}

func (helper *filePathHelper) GetFileName() string {
	slittedPath := strings.Split(*helper._filePathString, "/")
	return slittedPath[len(slittedPath)-1]
}

func (helper *filePathHelper) GetParentAbsolutePath() string {
	index := strings.LastIndex(*helper._filePathString, "/")
	return strings.Join(strings.Fields(*helper._filePathString)[0:index], "")
}

func (helper *filePathHelper) GetChildItems() *list.List {
	return list.New()
}

func (helper *filePathHelper) AppendPathToExisting(additionalPath *string) {
	//helper._filePathString = append(strings.Fields(*helper._filePathString), *additionalPath)
}
