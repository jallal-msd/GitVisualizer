package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func RecursiveScan(folder string) []string {

	return ScanGitFolders(make([]string, 0), folder)
}

func ScanGitFolders(folders []string, folder string) []string {

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
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = ScanGitFolders(folders, path)
		}
	}

	return folders

}
