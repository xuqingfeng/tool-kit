package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var (
	test_data = "test_data"
)

func TestCreateFile(t *testing.T) {

	var testFile = test_data + "/test"
	err := CreateFile(testFile)
	if err != nil {
		t.Fatalf("E! CreateFile() failed: %v", err)
	}
	_, err = os.Stat(testFile)
	if err != nil {
		t.Fatalf("E! can't get file status: %v", err)
	}

	defer func() {
		// cleanup
		os.Remove(testFile)
	}()
}

func TestCreateFileWithContent(t *testing.T) {

	var testFile = test_data + "/test"
	var testContent = "asdf"
	err := CreateFileWithContent(testFile, testContent)
	if err != nil {
		t.Fatalf("E! CreateFileWithContent() failed: %v", err)
	}

	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("E! read file failed: %v", err)
	}

	if testContent != fmt.Sprintf("%s", content) {
		t.Fatalf("E! file content don't match")
	}

	defer func() {
		// cleanup
		os.Remove(testFile)
	}()
}

func TestCreateDir(t *testing.T) {

	var testDir = test_data + "/test"
	err := CreateDir(testDir)
	if err != nil {
		t.Fatalf("E! CreateDir() failed: %v", err)
	}

	dir, err := os.Stat(testDir)
	if err != nil {
		t.Fatalf("E! can't get dir status: %v", err)
	}

	if !dir.IsDir() {
		t.Fatalf("E! %s is not a dir", testDir)
	}

	defer func() {
		// cleanup
		os.Remove(testDir)
	}()
}

func TestCreateSoftLink(t *testing.T) {

	var testFile = test_data + "/test"
	var testLink = test_data + "/test_link"

	CreateFile(testFile)
	err := CreateSoftLink(testFile, testLink)
	if err != nil {
		t.Fatalf("E! CreateSoftLink() failed: %v", err)
	}
	link, err := os.Lstat(testLink)
	if err != nil {
		t.Fatalf("E! can't get link status: %v", err)
	}
	if link.Mode()&os.ModeSymlink == 0 {
		t.Fatalf("E! %s is not a soft link", link.Name())
	}

	defer func() {
		// cleanup
		os.Remove(testFile)
		os.Remove(testLink)
	}()
}
