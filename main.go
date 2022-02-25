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
	fmt.Printf(bold + "Link Name: " + normal + "")
	// reading the String till I don't receive an enter ...
	l.lname, _ = reader.ReadString('\n')
	// trim \n from the end ...
	l.lname = strings.TrimSuffix(l.lname, "\n")

	fmt.Printf(bold + "Link: " + normal + "")
	l.link, _ = reader.ReadString('\n')
	l.link = strings.TrimSuffix(l.link, "\n")

	fmt.Printf(bold + "Description: " + normal + "")
	l.desc, _ = reader.ReadString('\n')
	l.desc = strings.TrimSuffix(l.desc, "\n")

	fmt.Printf(bold + "Category: " + normal + "")
	l.cate, _ = reader.ReadString('\n')
	l.cate = strings.TrimSuffix(l.cate, "\n")
}

func (l *Linker) generate_md() {
	fmt.Printf("\n* [**%s**](%s):\n\n**Description:**\n\n> %s\n\n**Category:** `%s`\n", l.lname, l.link, l.desc, l.cate)
}

// main function
func main() {
	myLinker := Linker{}
	myLinker.get_data()
	myLinker.generate_md()

}
