package cmd

import (
	"os"
	"testing"
)

func Test_GetDirPath(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	var vals = []struct {
		name     string
		inputVal []string
		expected string
	}{
		{"Accepts provided argument", []string{"/Users/bigsexy/Desktop/current/ulldSandbox/ulldApp"}, "/Users/bigsexy/Desktop/current/ulldSandbox/ulldApp"},
		{"Implements default CWD", []string{}, cwd},
	}
	for _, tt := range vals {
		t.Run(tt.name, func(t *testing.T) {
			res := GetDirPath(tt.inputVal)
			if res != tt.expected {
				t.Errorf("Expected %s got %s", tt.expected, res)
			}
		})
	}
}
