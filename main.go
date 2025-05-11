package main

import (
	"flag"
	"fmt"
)

func scan(folder string) {
	fmt.Println("starting scan..")
	repositories := RecursiveScan(folder)
	fmt.Println("get dotfile")
	filePath := getDotFilePath()
	fmt.Println("add slice to file")
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added \n\n")

}
func stats(path string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for git repo")
	flag.StringVar(&email, "email", "", "email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)

}
