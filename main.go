package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Println("starting scan..")

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
