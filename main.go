package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// NOTE: ALL USER SETTINGS !!
var repos = os.Getenv("REPOS")
var link_repo = repos + "links/"
var link_file = link_repo + "README.md"

// declaring constants for bold and normal text ...
const bold, normal string = "\033[1m", "\033[0m"

type Linker struct {
	lname, link, desc, cate, md string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func (l *Linker) get_data() {
	// creating a buffer to read input which comes from OS...
	reader := bufio.NewReader(os.Stdin)
	// Printing a formated string to take input ...
	fmt.Printf(bold + "LINK NAME: " + normal + "")
	// reading the String till I don't receive an enter ...
	l.lname, _ = reader.ReadString('\n')
	// trim \n from the end ...
	l.lname = strings.TrimSuffix(l.lname, "\n")

	fmt.Printf(bold + "LINK: " + normal + "")
	l.link, _ = reader.ReadString('\n')
	l.link = strings.TrimSuffix(l.link, "\n")

	fmt.Printf(bold + "DESCRIPTION: " + normal + "")
	l.desc, _ = reader.ReadString('\n')
	l.desc = strings.TrimSuffix(l.desc, "\n")

	fmt.Printf(bold + "CATEGORY: " + normal + "")
	l.cate, _ = reader.ReadString('\n')
	l.cate = strings.TrimSuffix(l.cate, "\n")

}

func (l *Linker) generate_md() string {
	// Manipulating Strings ... To get desired style...
	// Link Name: {Capitalize String}
	// Link: {Lowercase String}
	// Description: {Capitalize String}
	// Category: {UPPERCASE String}

	l.lname = strings.Title(l.lname)
	l.link = strings.ToLower(l.link)
	l.desc = strings.ToTitle(l.desc)
	l.cate = strings.ToUpper(l.cate)

	out := fmt.Sprintf("\n[**%s**](%s):\n", l.lname, l.link)
	out += fmt.Sprintf("\n> %s\n\n", l.desc)
	out += fmt.Sprintf("**Category:** `%s`\n", l.cate)
	return out
}

func (l *Linker) add_md() {

	// mkdir if not available

	// checks if testlink.md exists or not...
	// if not then create
	err := os.MkdirAll(link_repo, 0750)
	check(err)
	filehandler, err := os.OpenFile(link_file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	data := l.md
	if _, err := filehandler.WriteString(data); err != nil {
		check(err)
	}
	filehandler.Sync()

}

func show_config() {
	fmt.Println(bold + "Home Repository: " + normal + repos)
	fmt.Println(bold + "Link Repository: " + normal + link_repo)
	fmt.Println(bold + "Link Data File:  " + normal + link_file)
}

// main function
func main() {
	myLinker := Linker{}
	myLinker.get_data()
	myLinker.md = myLinker.generate_md()
	myLinker.add_md()
}
