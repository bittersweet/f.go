package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func visit(path string, f os.FileInfo, err error) error {
	red := "\033[31m"
	reset := "\033[0m"

	if path == "." {
		return nil
	}

	// Ignore .git and tmp directories
	ignore, _ := regexp.Compile(".git|tmp")
	if ignore.MatchString(path) == true {
		return nil
	}

	r, err := regexp.Compile(flag.Arg(0))

	if err != nil {
		fmt.Printf("Problem with regex: %v\n", err)
		return nil
	}

	if r.MatchString(path) == true {
		replaceWith := fmt.Sprintf("%v%v%v", red, flag.Arg(0), reset)
		formattedResult := r.ReplaceAllString(path, replaceWith)
		fmt.Println(formattedResult)
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
