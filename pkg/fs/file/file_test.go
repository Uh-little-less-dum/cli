package file

import (
	"testing"
)


const filePath string = "/Users/bigsexy/Desktop/Go/projects/ulld/fs/file/testFileContent.txt"
const dirName string = "/Users/bigsexy/Desktop/Go/projects/ulld/fs/file"


func getTestFile() File {
    return File{filePath}
}


func TestFileReadText(t *testing.T) {
	fileContentString := "Some text file content here.\n"
    file := getTestFile()
    fileContent := file.ReadText()
	if fileContent != fileContentString {
		t.Error("File did not return proper data from ReadText method.")
	}
}


func TestFileExists(t *testing.T) {
    file := getTestFile()
    if !file.Exists() {
        t.Error("File returns exists=false for a file that exists.")
    }

    fileNotExists := File{"/Users/bigsexy/Desktop/Go/projects/ulld/fs/file/path/does/not/exist/here/shouldfail.css"}

    if fileNotExists.Exists() {
        t.Error("File returns exists = true for file that does not exist.")
    }
}


func TestFileDirname(t *testing.T) {
    f := getTestFile()

    if f.Dirname() != dirName {
        t.Error("File does not return proper dirname")
    }
}


func TestGetParentDir(t *testing.T) {
    file := getTestFile()
    dir := file.GetParentDir()
    if dir.Path() != dirName {
        t.Error("File GetParentDir method does not return a directory struct with the proper path.")
    }
}
