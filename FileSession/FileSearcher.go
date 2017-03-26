package FileSession

// IFileSearcher searcher for files in directory
type IFileSearcher interface {
}

type fileSearcher struct {
}

// FileSearcherFactory creates and inits an instance of IFileSearcher
func FileSearcherFactory() IFileSearcher {
	fileSearcherInstance := new(fileSearcher)
	return fileSearcherInstance
}
