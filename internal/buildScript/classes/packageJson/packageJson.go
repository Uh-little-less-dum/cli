package package_json

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
)

type PackageJsonSchema interface{}

// RESUME: Come back here and generate the schema with quicktype when back online.
type PackageJsonFile struct {
	path string
	data PackageJsonSchema
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (f PackageJsonFile) Path() string {
	return f.path
}

func (f PackageJsonFile) SetPath(newPath string) error {
	f.path = newPath
	return f.Read()
}

func (f *PackageJsonFile) Read() error {
	b, err := os.ReadFile(f.path)
	if err != nil {
		return err
	}
	if !json.Valid(b) {
		log.Fatalf("The json at %s is not valid json. We can't continue", f.path)
	}
	f.data = string(b)
	return nil
}

// RESUME: Move this to the developer cli and implement this gjson functionality when able to look at the docs again.
// func (f PackageJsonFile) Query(dataPath string) any {
//     return gjson.Result
// }
