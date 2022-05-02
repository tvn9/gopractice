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

import (
	"fmt"
	"log"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func main() {
	// Server list
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	// Define srvStat map variable
	srvStat := make(map[string]int)

	// Insert servers name and set status to 0 to map
	for _, v := range servers {
		srvStat[v] = Online
	}

	// Print Server stat
	serverStat(srvStat)

	//  - change server status of `darkstar` to `Retired`
	fmt.Println("\nAfter changed darkstar status")
	srvStat["darkstar"] = Retired
	serverStat(srvStat)

	//  - change server status of `aiur` to `Offline`
	srvStat["aiur"] = Offline
	fmt.Println("\nAfter changed aiur status")
	serverStat(srvStat)

	//  - change server status of all servers to `Maintenance`
	for k := range srvStat {
		srvStat[k] = Maintenance
	}
	fmt.Println("\nAfter changed all servers to Maintenance")
	serverStat((srvStat))

}

func serverStat(s map[string]int) {

	// Print total number of server
	fmt.Println("\nNumber of server =", len(s))

	var (
		nOnline, nOffline, nMaintenance, nRetired int
		stat                                      string
	)

	// Print details server status
	for k, v := range s {

		switch v {
		case Online:
			nOnline += 1
			stat = "online"
			fmt.Printf("Server %q status: %s\n", k, stat)
		case Offline:
			nOffline += 1
			stat = "Offline"
			fmt.Printf("Server %q status: %s\n", k, stat)
		case Maintenance:
			nMaintenance += 1
			stat = "Maintenance"
			fmt.Printf("Server %q status: %s\n", k, stat)
		case Retired:
			nRetired += 1
			stat = "Retired"
			fmt.Printf("Server %q status: %s\n", k, stat)
		default:
			log.Fatal("Exit 1")
		}
	}
	fmt.Println("Number of online server:", nOnline)
	fmt.Println("Number of offline server:", nOffline)
	fmt.Println("Number of maintainance server:", nMaintenance)
	fmt.Println("Number of retired server:", nRetired)

}
