package manage_dir

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func AskFilePrefix() string {
	// Ask for a input that will be used to rename files
	// Return a string
	fmt.Println("Provide a prefix for your files")

	var prefix string

	fmt.Scanln(&prefix)

	return prefix
}

func scanCurrentDir(prefix string) ([]string, int) {
	directory := "."
	var filesNames []string
	var start int
	programName := filepath.Base(os.Args[0])
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Dir(path) == directory && info.Name() != programName && !strings.HasPrefix(info.Name(), prefix) {
			filesNames = append(filesNames, info.Name())
		}
		if strings.HasPrefix(info.Name(), prefix) {
			parts := strings.Split(info.Name(), "_")
			if len(parts) != 2 {
				fmt.Println("Wrong format !")
				return nil
			}
			subparts := strings.Split(parts[1], ".")
			if len(subparts) != 2 {
				fmt.Println("Wrong format !")
				return nil
			}
			numericPart := subparts[0]
			number, err := strconv.Atoi(numericPart)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return err
			}
			start = number
		}
		return err
	})

	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	return filesNames, start
}
