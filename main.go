package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// declaring constants for bold and normal text ...
const bold, normal string = "\033[1m", "\033[0m"

type Linker struct {
	lname, link, desc, cate string
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

	l.lname = strings.ToTitle(l.lname)
	l.link = strings.ToLower(l.link)
	l.desc = strings.ToTitle(l.desc)
	l.cate = strings.ToUpper(l.cate)

	out := fmt.Sprintf("\n* [**%s**](%s):\n", l.lname, l.link)
	out += fmt.Sprintf("\n**Description:**\n\n> %s\n\n", l.desc)
	out += fmt.Sprintf("**Category:** `%s`\n", l.cate)
	return out
}

// main function
func main() {
	myLinker := Linker{}
	myLinker.get_data()
	fmt.Println(myLinker.generate_md())
}
