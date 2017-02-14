package MachineLearning

import (
	"strings"
)

// FileType type of file where to save
type FileType uint32

// Type of files subdivs
const (
	Undefined    FileType = 0
	Image        FileType = 1
	Literature   FileType = 2
	Documents    FileType = 3
	Video        FileType = 4
	Presentation FileType = 5
	SourceCode   FileType = 6
)

// IParseFileEnding parse file and return where it belongs
type IParseFileEnding interface {
	ParseFile(filePath *string) FileType
}

// ParseFile parses file and return type based on prefix
func ParseFile(filePath *string) FileType {
	fileType := Undefined
	if strings.HasPrefix(*filePath, ".png") || strings.HasPrefix(*filePath, ".bmp") ||
		strings.HasPrefix(*filePath, ".gif") {
		fileType = Image
	}

	if strings.HasPrefix(*filePath, ".pptx") {
		fileType = Presentation
	}

	if strings.HasPrefix(*filePath, ".pdf") {
		fileType = Literature
	}

	if strings.HasPrefix(*filePath, ".vma") || strings.HasPrefix(*filePath, ".mp4") ||
		strings.HasPrefix(*filePath, ".mpeg") {
		fileType = Video
	}

	if strings.HasPrefix(*filePath, ".docx") || strings.HasPrefix(*filePath, ".doc") ||
		strings.HasPrefix(*filePath, ".txt") {
		fileType = Documents
	}

	if strings.HasPrefix(*filePath, ".cpp") || strings.HasPrefix(*filePath, ".go") ||
		strings.HasPrefix(*filePath, ".js") || strings.HasPrefix(*filePath, ".h") || strings.HasPrefix(*filePath, ".cs") ||
		strings.HasPrefix(*filePath, ".java") || strings.HasPrefix(*filePath, ".py") {
		fileType = SourceCode
	}

	return fileType
}
