//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func display(s map[string]int) {
	//* Create a function to print server status displaying:
	fmt.Println(s)
	//  - number of servers
	//  - number of servers for each status (Online, Offline, Maintenance, Retired)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	myMap := make(map[string]int)
	for _, status := range servers {
		myMap[status] = Online
	}
	display(myMap)
	myMap["darkstar"] = Retired
	display(myMap)
	myMap["aiur"] = Offline
	display(myMap)
	for _, status := range servers {
		myMap[status] = Maintenance
	}
	display(myMap)
	// 	After creating the map, perform the following actions:
	//  - call display server info function
	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	//  - call display server info function
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
}
