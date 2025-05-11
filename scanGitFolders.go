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

func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dotFile := usr.HomeDir + "/.gogitlocalstats"
	return dotFile

}

func addNewSliceElementsToFile(filepath string, repo []string) {

	existingRepos := parseFileLinesToSlice(filepath)
	repos := joinSlices(repo, existingRepos)
	dumbStringSliceToFile(repos, filepath)

}
func parseFileLinesToSlice(filepath string) []string {
	f := openFile(filepath)
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

func openFile(filepath string) *os.File {

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			panic(err)
		}
	}
	return f
}
func joinSlices(neew []string, existing []string) []string {

	for _, i := range neew {
		if !sliceContaines(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing

}
func sliceContaines(slice []string, value string) bool {

	for _, val := range slice {
		if val == value {
			return true

		}
	}
	return false
}

func dumbStringSliceToFile(repos []string, filePath string) {

	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filePath, []byte(content), 0755)

}
