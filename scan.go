package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
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

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotfilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "./gogitlocalstats"

	return dotFile
}

func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}

func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}

	return existing
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func dumpStringSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filePath, []byte(content), 0755)
}

func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringSliceToFile(repos, filePath)
}

func scan(folder string) {
	fmt.Printf("Found folders:\n\n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotfilePath()
	addNewSliceElementsToFile(filepath, repositories)
	fmt.Printf("\nSuccessfully added\n\n")
}
