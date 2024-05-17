//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

//* Create a structure to store items and their security tag state
type Item struct {
	Name string
	tag  SecurityTag
}

func activateTag(tag *SecurityTag) {
	*tag = Active
}

func deactivateTag(tag *SecurityTag) {
	*tag = Inactive
}

// Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	for i := range items {
		deactivateTag(&items[i].tag)
	}
}

func main() {
	mySlice := []Item{{"item1", Active}, {"item2", Active}, {"item3", Active}, {"item4", Active}}
	fmt.Println(mySlice)
	deactivateTag(&mySlice[2].tag)
	fmt.Println(mySlice)
	checkout(mySlice)
	fmt.Println(mySlice)

}
