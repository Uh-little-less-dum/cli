package utils_error

import "github.com/charmbracelet/log"

// TODO: Remove this entire file. This was moved to the goutils package.
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
