package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	root := "/home/tamal/go/src/kubeform.dev/installer/.generator/hack"

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		fmt.Println(rel)

		return nil
	})
	if err != nil {
		panic(err)
	}
}
