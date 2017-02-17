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

}

func (helper *filePathHelper) IsAbsolutePath() bool {

}

func (helper *filePathHelper) IsRelativePath() bool {

}

func (helper *filePathHelper) IsFilePathExists() bool {

}

func (helper *filePathHelper) GetRelativePath() string {

}

func (helper *filePathHelper) GetAboslutePath() string {

}

func (helper *filePathHelper) GetParentName() string {

}

func (helper *filePathHelper) GetFileName() string {

}

func (helper *filePathHelper) GetParentAbsolutePath() string {

}

func (helper *filePathHelper) GetChildItems() *list.List {

}

func (helper *filePathHelper) AppendPathToExisting(additionalPath *string) {

}
