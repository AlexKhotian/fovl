package MachineLearning

import (
	"log"
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
	Media        FileType = 4
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
	if strings.HasSuffix(*filePath, ".png") || strings.HasSuffix(*filePath, ".bmp") ||
		strings.HasSuffix(*filePath, ".gif") {
		fileType = Image
	}

	if strings.HasSuffix(*filePath, ".pptx") {
		fileType = Presentation
	}

	if strings.HasSuffix(*filePath, ".pdf") {
		fileType = Literature
	}

	if strings.HasSuffix(*filePath, ".vma") || strings.HasSuffix(*filePath, ".mp4") ||
		strings.HasSuffix(*filePath, ".mpeg") || strings.HasSuffix(*filePath, ".mp3") ||
		strings.HasSuffix(*filePath, ".avi") {
		fileType = Media
	}

	if strings.HasSuffix(*filePath, ".docx") || strings.HasSuffix(*filePath, ".doc") ||
		strings.HasSuffix(*filePath, ".txt") {
		fileType = Documents
	}

	if strings.HasSuffix(*filePath, ".cpp") || strings.HasSuffix(*filePath, ".go") ||
		strings.HasSuffix(*filePath, ".js") || strings.HasSuffix(*filePath, ".h") || strings.HasSuffix(*filePath, ".cs") ||
		strings.HasSuffix(*filePath, ".java") || strings.HasSuffix(*filePath, ".py") {
		fileType = SourceCode
	}

	if fileType == Undefined {
		log.Println("ParseFile: Failed to obtain filetype for file: ", filePath)
		fileType = Documents
	}

	return fileType
}
