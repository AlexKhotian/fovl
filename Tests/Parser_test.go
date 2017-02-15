package Tests

import "testing"
import "smart_file_transport/MachineLearning"

func TestParsePresentation(t *testing.T) {
	testString := "test.pptx"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Presentation {
		t.Log("Failed to parse Presentation")
		t.Fail()
	}
}

func TestParseImage(t *testing.T) {
	testString := "test.png"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Image {
		t.Log("Failed to parse image")
		t.Fail()
	}
}

func TestParseDocument(t *testing.T) {
	testString := "test.docx"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Documents {
		t.Log("Failed to parse Documents")
		t.Fail()
	}
}

func TestParseLiterature(t *testing.T) {
	testString := "test.pdf"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Literature {
		t.Log("Failed to parse Literature")
		t.Fail()
	}
}

func TestParseSourceCode(t *testing.T) {
	testString := "test.cpp"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.SourceCode {
		t.Log("Failed to parse SourceCode")
		t.Fail()
	}
}

func TestParseFilesMedia(t *testing.T) {
	testString := "test.mp4"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Media {
		t.Log("Failed to parse Media")
		t.Fail()
	}
}
