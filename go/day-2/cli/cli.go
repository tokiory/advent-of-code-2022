package cli

import "os"

func GetPath() string {
	if len(os.Args) < 2 {
		panic("Specify the path for input file")
	}

	return os.Args[1]
}
