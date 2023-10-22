package main

import (
	// read_dir handle the files renaming
	"change_files_name/pkg/read_dir"
)

func main() {
	// Ask for the prefix
	my_prefix := read_dir.AskFilePrefix()

	// Rename every files in the current folder
	read_dir.RenameFiles(my_prefix)
}
