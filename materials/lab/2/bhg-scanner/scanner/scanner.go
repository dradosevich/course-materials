// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {in the main directory 'go build main' './main'}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil {
			results <- -1 * p //this seemed to be the most elligant solution
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(start int, stop int) (int, int) {
	//TODO 3 : ADD closed ports; currently code only tracks open ports
	var openports []int          // notice the capitalization here. access limited!
	var closedports []int        //still not exported
	ports := make(chan int, 200) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := start; i <= stop; i++ {
			ports <- i
		}
	}()

	for i := start; i <= stop; i++ {
		port := <-results
		if port > 0 {
			openports = append(openports, port)
		} else if port < 0 {
			closedports = append(closedports, -1*port)
		}
	}

	defer close(ports)
	defer close(results)
	sort.Ints(openports)
	sort.Ints(closedports) //sort our closed ports too

	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for i, port := range openports {
		if i%5 == 0 {
			fmt.Println()
		}
		fmt.Printf("%04d:open\t,\t", port)
	}
	for i, port := range closedports {
		if i%5 == 0 {
			fmt.Println()
		}
		fmt.Printf("%04d:closed\t,\t", port)
	}
	fmt.Println()
	fmt.Println()

	return len(openports), len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed);
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
