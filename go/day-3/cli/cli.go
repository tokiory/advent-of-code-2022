package cli

import "os"

func GetPath() string {
	if len(os.Args) < 2 {
		panic("File for parsing doesn't specified")
	}

	return os.Args[1]
}
