package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "./git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
		}
		if file.Name() == "vendor" || file.Name() == "node_modules" {
			continue
		}
		folders = scanGitFolders(folders, path)
	}

	return folders
}

func scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotfilePath()
	addNewSliceElementsToFile(filepath, repositories)
	fmt.Printf("\nSuccessfully added\n\n")
}
