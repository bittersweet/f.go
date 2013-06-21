package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func visit(path string, f os.FileInfo, err error) error {
	if path == "." {
		return nil
	}

	r, err := regexp.Compile(flag.Arg(0))

	if err != nil {
		fmt.Printf("Problem with regex: %v\n", err)
		return nil
	}

	if r.MatchString(path) == true {
		fmt.Println(path)
	}

	return nil
}

func main() {
	flag.Parse()
	root := "."
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("returned %v\n", err)
	}
}
