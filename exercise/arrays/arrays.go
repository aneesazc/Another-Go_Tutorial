//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Item struct {
	prodName string
	price    int
}

func main() {
	myArr := make([]Item, 4)
	myArr = append(myArr, Item{"cup", 45})
	myArr = append(myArr, Item{"bat", 100})
	myArr = append(myArr, Item{"ball", 35})
	fmt.Println(myArr[len(myArr)-1])
	fmt.Println(len(myArr))
	cost := 0
	for i := 0; i < len(myArr); i++ {
		cost += myArr[i].price
	}
	fmt.Println(cost)
	myArr = append(myArr, Item{"shirt", 125})

}
