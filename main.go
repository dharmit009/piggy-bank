package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// NOTE: ALL USER SETTINGS !!
var repos = os.Getenv("REPOS")
var link_repo = repos + "links/"
var link_file = link_repo + "README.md"
var json_file = link_repo + "data.json"

// declaring constants for bold and normal text ...
const bold, normal string = "\033[1m", "\033[0m"

type linker struct {
	Lname, Link, Desc, Cate, Md string
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func (l *linker) get_data() {

	/** creating a buffer to read input which comes from OS...
	Printing a formated string to take input ...
	reading the String till I don't receive an enter ...
	trim \n from the end ... **/

	myreader := bufio.NewReader(os.Stdin)
	fmt.Printf(bold + "LINK NAME: " + normal + "")
	l.Lname, _ = myreader.ReadString('\n')

	fmt.Printf(bold + "LINK: " + normal + "")
	l.Link, _ = myreader.ReadString('\n')

	fmt.Printf(bold + "DESCRIPTION: " + normal + "")
	l.Desc, _ = myreader.ReadString('\n')

	fmt.Printf(bold + "CATEGORY: " + normal + "")
	l.Cate, _ = myreader.ReadString('\n')

	l.Lname = strings.TrimSuffix(l.Lname, "\n")
	l.Link = strings.TrimSuffix(l.Link, "\n")
	l.Desc = strings.TrimSuffix(l.Desc, "\n")
	l.Cate = strings.TrimSuffix(l.Cate, "\n")

}

func (l *linker) generate_md() string {

	// Manipulating Strings ... To get desired style...
	// Link Name: {Capitalize String}
	// Link: {Lowercase String}
	// Description: {Capitalize String}
	// Category: {UPPERCASE String}

	l.Lname = strings.Title(l.Lname)
	l.Link = strings.ToLower(l.Link)
	l.Desc = strings.ToTitle(l.Desc)
	l.Cate = strings.ToUpper(l.Cate)

	out := fmt.Sprintf("\n[**%s**](%s):\n", l.Lname, l.Link)
	out += fmt.Sprintf("\n> %s\n\n", l.Desc)
	out += fmt.Sprintf("**Category:** `%s`\n", l.Cate)
	return out
}

func (l *linker) add_md() {

	/** mkdir if not available
	checks if testlink.md exists or not...
	if not then create **/
	err := os.MkdirAll(link_repo, 0750)
	check(err)
	filehandler, err := os.OpenFile(link_file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	data := l.Md
	if _, err := filehandler.WriteString(data); err != nil {
		check(err)
	}
	filehandler.Sync()

}

func (l *linker) encode_json() {
	indata := []linker{
		{l.Lname, l.Link, l.Desc, l.Cate, l.Md},
	}
	final_json, _ := json.MarshalIndent(indata, "", "\t")
	fmt.Println(string(final_json))

	err := os.MkdirAll(link_repo, 0750)
	check(err)
	jsonhandler, err := os.OpenFile(json_file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	data := string(final_json)
	if _, err := jsonhandler.WriteString(data); err != nil {
		check(err)
	}
	jsonhandler.Sync()
}

func show_config() {
	fmt.Println(bold + "Home Repository: " + normal + repos)
	fmt.Println(bold + "Link Repository: " + normal + link_repo)
	fmt.Println(bold + "Link Data File:  " + normal + link_file)
	fmt.Println(bold + "Json Data File:  " + normal + json_file)
} //show_config...

func committer() {

}

// main function
func main() {
	my_linker := linker{}
	my_linker.get_data()
	my_linker.Md = my_linker.generate_md()
	my_linker.add_md()
	my_linker.encode_json()
}
