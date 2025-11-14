// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-14
// This program calculates the cost of bolts, nuts, and washers,
// checks if the order is valid, and displays the total cost.

package main

import (
	"fmt"
)

func main() {

	// Constants
	const BOLT_COST = 10   // cents
	const NUT_COST = 5     // cents
	const WASHER_COST = 3  // cents

	var bolts int
	var nuts int
	var washers int

	// Inputs
	fmt.Print("How many bolts would you like to purchase? ")
	fmt.Scan(&bolts)

	fmt.Print("How many nuts would you like to purchase? ")
	fmt.Scan(&nuts)

	fmt.Print("How many washers would you like to purchase? ")
	fmt.Scan(&washers)

	// Output counts
	fmt.Println("Number of bolts:          ", bolts)
	fmt.Println("Number of nuts:           ", nuts)
	fmt.Println("Number of washers:        ", washers)

	// Order checks
	if nuts < bolts {
		fmt.Println("Check the Order - not enough nuts for the bolts you purchased.")
	}

	if washers < bolts {
		fmt.Println("Check the Order - not enough washers for the bolts you purchased.")
	}

	if nuts >= bolts && washers >= bolts {
		fmt.Println("Order is OK.")
	}

	// Total cost
	totalCost := bolts*BOLT_COST + nuts*NUT_COST + washers*WASHER_COST
	fmt.Printf("Your total cost of the order is %d cents.\n", totalCost)

	// Done message
	fmt.Println("\nDone.")
}

