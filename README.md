# Usage

	package main

	import (
		"fmt"
		"github.com/freehaha/ignores"
		"os"
		"path/filepath"
	)

	var ig *ignores.Ignore

	func main() {
		ig, _ = ignores.New(".ignore")
		filepath.Walk(".", walkFunc)
	}

	func walkFunc(path string, info os.FileInfo, err error) error {
		ignore, err := ig.Match(path)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		if ignore {
			if info.IsDir() {
				fmt.Printf("ignored: %s\n", path)
				return filepath.SkipDir
			}
			fmt.Printf("ignored: %s\n", path)
			return nil
		}
		fmt.Printf("path: %s\n", path)
		return nil
	}
