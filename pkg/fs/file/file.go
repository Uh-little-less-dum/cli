package file

import (
	"bufio"
	"errors"
	"github.com/igloo1505/ulldCli/pkg/fs/directory"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type File struct {
	path string
}

func (f File) Path() string {
	return f.path
}

func (f File) Dirname() string {
	return filepath.Dir(f.path)
}

func (f File) Join(filePaths ...string) string {
	allFilePaths := []string{f.path}
	allFilePaths = append(allFilePaths, filePaths...)
	return filepath.Join(allFilePaths...)
}

// Don't forget to close the file!
func (f File) GetOsFile() *os.File {
	fileData, err := os.Open(f.path)
	check(err)
	return fileData
}

// Don't forget to close the file!
func (f File) GetReader() (*os.File, *bufio.Reader) {
	fileData := f.GetOsFile()
	reader := bufio.NewReader(fileData)
	return fileData, reader
}

func (f File) Exists() bool {
	_, err := os.Stat(f.path)
	return !errors.Is(err, os.ErrNotExist)
}

func (f File) CreateIfNotExists() {
	if !f.Exists() {
		_, err := os.Create(f.path)
		check(err)
	}
}

func (f File) ReadText() string {
	data, err := os.ReadFile(f.path)
	check(err)
	return string(data)
}

func (f File) WriteText(content string) {
	osFile := f.GetOsFile()
	_, err := osFile.WriteString(content)
	check(err)
	err = osFile.Sync()
	check(err)
	osFile.Close()
}

func (f File) GetParentDir() directory.Directory {
	dirname := f.Dirname()
	var dir = directory.Directory{}
	dir.SetPath(dirname)
	return dir
}
