package utils_error

import "github.com/charmbracelet/log"

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
