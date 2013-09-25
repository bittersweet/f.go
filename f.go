package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func visit(path string, file os.FileInfo, err error) error {
	query := flag.Arg(0)

	if file.IsDir() {
		return nil
	}

	// Ignore .git directory
	ignore, _ := regexp.Compile("\\.git")
	if ignore.MatchString(path) == true {
		return nil
	}

	r, err := regexp.Compile(query)
	if err != nil {
		fmt.Printf("Problem with regex: %v\n", err)
		return nil
	}

	if r.MatchString(path) == true {
		replaceWith := colorizeMatch(query)
		formattedResult := r.ReplaceAllString(path, replaceWith)
		fmt.Println(formattedResult)
	}

	return nil
}

func colorizeMatch(match string) string {
	red := "\033[31m"
	reset := "\033[0m"

	return fmt.Sprintf("%v%v%v", red, match, reset)
}

func main() {
	flag.Parse()
	root := "."
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("returned %v\n", err)
	}
}
