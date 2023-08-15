package main

import (
	"change_files_name/pkg/read_dir"
)

func main() {
	my_prefix := read_dir.AskFilePrefix()
	read_dir.RenameFiles(my_prefix)
}
