package Tests

import "testing"
import "smart_file_transport/MachineLearning"

func TestParseFiles(t *testing.T) {
	testString := "test.png"
	fileType := MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Image {
		t.Log("Failed to parse image")
		t.Fail()
	}

	testString = "test.pdf"
	fileType = MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Literature {
		t.Log("Failed to parse Literature")
		t.Fail()
	}

	testString = "test.docx"
	fileType = MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Documents {
		t.Log("Failed to parse Documents")
		t.Fail()
	}

	testString = "test.mp4"
	fileType = MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Video {
		t.Log("Failed to parse Video")
		t.Fail()
	}

	testString = "test.pptx"
	fileType = MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.Presentation {
		t.Log("Failed to parse Presentation")
		t.Fail()
	}

	testString = "test.cpp"
	fileType = MachineLearning.ParseFile(&testString)
	if fileType != MachineLearning.SourceCode {
		t.Log("Failed to parse SourceCode")
		t.Fail()
	}
}
