//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string // member name

type LendAudit struct {
	checkOutTime time.Time
	checkInTime  time.Time // if this is zero value, then we know the book is still checked out and the member has not returned it
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int // total books lended out
}

type Library struct {
	books   map[Title]BookEntry
	members map[Name]Member // library will track members and members will have LentAudit checks
}

func printMemeberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkInTime.IsZero() {
			returnTime = "[not returned]"
		} else {
			returnTime = audit.checkInTime.String()
		}
		fmt.Printf("Book: %s, Check out time: %s, Check in time: %s\n", title, audit.checkOutTime.String(), returnTime)
	}
}

func printLibraryInfo(library *Library) {
	fmt.Println("Library Info:")
	for _, member := range library.members {
		fmt.Printf("Member: %s\n", member.name)
		printMemeberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println("Library Books:")
	for title, entry := range library.books {
		fmt.Printf("Book: %s, Total: %d, Lended: %d\n", title, entry.total, entry.lended)
	}
}

// func for checkout a book
func checkOutBook(library *Library, member *Member, title Title) {
	if entry, ok := library.books[title]; ok {
		if entry.lended < entry.total {
			// book is available
			entry.lended++
			library.books[title] = entry
			member.books[title] = LendAudit{checkOutTime: time.Now()}
		} else {
			fmt.Println("Book is not available")
		}
	} else {
		fmt.Println("Book not found")
	}
}

// func for check in a book
func checkInBook(library *Library, member *Member, title Title) {
	if _, ok := member.books[title]; ok {
		delete(member.books, title)
		entry := library.books[title]
		entry.lended--
		library.books[title] = entry
	} else {
		fmt.Println("Book not found")
	}
}

func main() {
	library := Library{
		books: map[Title]BookEntry{
			"Book1": {total: 2},
			"Book2": {total: 1},
			"Book3": {total: 3},
			"Book4": {total: 1},
		},
		members: map[Name]Member{
			"Member1": {name: "Member1", books: map[Title]LendAudit{}},
			"Member2": {name: "Member2", books: map[Title]LendAudit{}},
			"Member3": {name: "Member3", books: map[Title]LendAudit{}},
		},
	}

	printLibraryInfo(&library)
	printLibraryBooks(&library)

	member1 := library.members["Member1"]
	checkOutBook(&library, &member1, "Book1")
	checkOutBook(&library, &member1, "Book2")
	checkOutBook(&library, &member1, "Book3")
	checkOutBook(&library, &member1, "Book4")
	checkOutBook(&library, &member1, "Book5")
	checkInBook(&library, &member1, "Book1")
	checkInBook(&library, &member1, "Book2")

}
