package Tests

import (
	"fovl/FileSession"
	"testing"
)

func TestIsRelativePath(t *testing.T) {
	testPath := "test"
	helper := FileSession.CreateFilePathHelper(&testPath)
	if !helper.IsRelativePath() {
		t.Log("Failed to retrive relative path")
		t.Fail()
	}
}

func TestIsAbsolutePath(t *testing.T) {
	testPath := "/home/test1/test"
	helper := FileSession.CreateFilePathHelper(&testPath)
	if !helper.IsAbsolutePath() {
		t.Log("Failed to retrive absolute path")
		t.Fail()
	}
}

func TestGetRelativePath(t *testing.T) {
	testPath := "/test1/test2/test3"
	checkPath := "test3"
	helper := FileSession.CreateFilePathHelper(&testPath)
	if helper.GetRelativePath() != checkPath {
		t.Log("Failed to retrive relative path")
		t.Fail()
	}
}

func TestGetAbsolutePath(t *testing.T) {
	testPath := "/test1/test2/test3"
	helper := FileSession.CreateFilePathHelper(&testPath)
	if helper.GetAboslutePath() != testPath {
		t.Log("Failed to retrive absolute path")
		t.Fail()
	}
}

func TestGetParentPath(t *testing.T) {
	testPath := "/test1/test2/test3"
	checkPath := "test2"
	helper := FileSession.CreateFilePathHelper(&testPath)
	if helper.GetParentName() != checkPath {
		t.Log("Failed to retrive parent name")
		t.Fail()
	}
}

func TestGetParentAbsolutePath(t *testing.T) {
	/*	testPath := "/test1/test2/test3"
		checkPath := "/test1/test2/"
		helper := FileSession.CreateFilePathHelper(&testPath)
		if helper.GetParentAbsolutePath() != checkPath {
			t.Log("Failed to retrive parent absolute path")
			t.Fail()
		}*/
}
