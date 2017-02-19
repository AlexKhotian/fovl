package FileSession

import (
	"container/list"
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
func CreateFilePathHelper() IFilePathHelper {
	this := new(filePathHelper)
	return this
}

func (helper *filePathHelper) IsDirectory() bool {
	return false
}

func (helper *filePathHelper) IsAbsolutePath() bool {
	return false
}

func (helper *filePathHelper) IsRelativePath() bool {
	return false
}

func (helper *filePathHelper) IsFilePathExists() bool {
	return false
}

func (helper *filePathHelper) GetRelativePath() string {
	return ""
}

func (helper *filePathHelper) GetAboslutePath() string {
	return ""
}

func (helper *filePathHelper) GetParentName() string {
	return ""
}

func (helper *filePathHelper) GetFileName() string {
	return ""
}

func (helper *filePathHelper) GetParentAbsolutePath() string {
	return ""
}

func (helper *filePathHelper) GetChildItems() *list.List {
	return list.New()
}

func (helper *filePathHelper) AppendPathToExisting(additionalPath *string) {

}
