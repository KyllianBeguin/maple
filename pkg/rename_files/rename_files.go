package read_dir

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type nameHandler struct {
	OldName string
	NewName string
}

func makeHandlers(prefix string) []nameHandler {
	my_files, starting_val := scanCurrentDir(prefix)
	Handlers := make([]nameHandler, len(my_files))
	for index, old_name := range my_files {
		number := index + 1 + starting_val
		new_name := fmt.Sprintf("%s_%02d", prefix, number)
		Handlers[index].OldName = old_name
		Handlers[index].NewName = new_name + filepath.Ext(old_name)
	}
	return Handlers
}

func RenameFiles(prefix string) {
	handlers := makeHandlers(prefix)

	for _, hdlr := range handlers {
		err := os.Rename(hdlr.OldName, hdlr.NewName)
		if err != nil {
			log.Fatalf("Error renaming file: %v", err)
		}
	}
	if len(handlers) == 0 {
		fmt.Println("No files to rename 😮")
	} else {
		fmt.Println("File renamed successfully 🚀")
	}
}
