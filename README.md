# Piggy Bank ...

Remember those childhood days when we use to collect coins within house,
take a few cents from mom, and store it into piggy bank? Slowly and
slowly we have save a good amount which we are proud of I am applying
the same thing for links. So the basic idea is to store different links
into different files based on their categories. This will be a CLI tool
which will allow us to add, remove, search, list links based on
categories, open links and display links (all or on based on
categories). That's all this does one thing that is manage your links !!

## USAGE:

bank <options> ...

## OPTIONS:

1. Add.
1. Remove.
1. Search.
1. Open Link.
1. Category (add or list nothing else).
1. Show Based On Category and Showall.

## CATEGORIES:

1. Programming.
1. Cybersecurity.
1. Webdev.
1. Data-science.
1. GitHub-Links.

### SPECIFICATIONS && FEATURES:

1. Add functionality to select category while taking input for the same.
1. All the links will be stored in markdown files (.md) and will be uploaded to
GitHub as a backup. We will use Common-mark as our Markdown flavour.

### Prototype:

Add() --> nameOfLink --> Link --> Description (this cannot be empty) -->
Choose Category or Add new category --> Append the Link and Link
Information --> Commit changes and push to github.

Remove() --> nameOfLink --> grep the link --> make sure you want to
remove the link [y/n] --> If removed save it in a trash category for
backup.

search() --> Take some few keywords --> pass that grep --> Find links
and display them.

open() --> Select Category --> Select Link --> Open it in default
browser.

show <category> and showall --> Shows all the links from all categories,
otherwise you can select or type a <category> to show all the links of
the respective category.


