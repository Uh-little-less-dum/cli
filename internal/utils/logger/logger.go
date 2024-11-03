package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func DebugLog(val ...string) {
	if len(os.Getenv("DEBUG")) <= 0 {
		return
	}

	file_path := "/Users/bigsexy/Desktop/Go/projects/ulld/cli/debug.log"

	content := fmt.Sprintf("%s\n", strings.Join(val, " "))
	_, err := os.Stat(file_path)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File doesn't exist")
	} else {
		fmt.Println(os.Getenv("UlldCliHasLogged"))
		if os.Getenv("UlldCliHasLogged") != "true" {
			f, err := os.OpenFile(file_path, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.WriteString("")
			if err != nil {
				cobra.CheckErr(err)
			}
			err = f.Sync()
			if err != nil {
				cobra.CheckErr(err)
			}
			f.Close()
		}
		f, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			cobra.CheckErr(err)
		} else {
			os.Setenv("UlldCliHasLogged", "true")
		}
		defer f.Close()

		if _, err := f.WriteString(content); err != nil {
			cobra.CheckErr(err)
		}
	}
}

// func DebugLog() {
// 	if len(os.Getenv("DEBUG")) > 0 {
// 		f, err := tea.LogToFile("debug.log", "debug")
// 		if err != nil {
// 			fmt.Println("fatal:", err)
// 			os.Exit(1)
// 		}
// 		defer f.Close()
// 	}
// }
