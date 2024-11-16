package directory

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Directory struct {
	path string
}

func (d Directory) Path() string {
	return d.path
}

func (d *Directory) SetPath(filePath string) {
	d.path = filePath
}

func (d Directory) Exists() bool {
	_, err := os.Stat(d.path)
	return !errors.Is(err, os.ErrNotExist)
}

func (d Directory) MkDirIfNotExists() {
	exists := d.Exists()
	if exists {
		return
	}
	err := os.MkdirAll(d.path, 0755)
	check(err)
}

// Change directory to the provided directory.
func (d Directory) Chdir(makeIfNotExists bool) {
	if makeIfNotExists {
		d.MkDirIfNotExists()
	}
	err := os.Chdir(d.path)
	check(err)
}

// Calls function for every child file and directory path recursively.
func (d Directory) Walk(f fs.WalkDirFunc) {
	err := filepath.WalkDir(d.path, f)
	check(err)
}

func (d Directory) GetDirContents() []fs.DirEntry {
	c, err := os.ReadDir(d.path)
	check(err)
	return c
}

func (d Directory) CreateEmptyFile(filePath string) {
	byteArr := []byte("")
	err := os.WriteFile(filePath, byteArr, 0644)
	check(err)
}
