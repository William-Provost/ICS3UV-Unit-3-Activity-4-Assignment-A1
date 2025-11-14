/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-14
 * @fileoverview This program calculates the cost of nuts, bolts, and washers,
 * and checks whether the order is valid.
 */

// Constants
const BOLT_COST = 10;    // cents
const NUT_COST = 5;      // cents
const WASHER_COST = 3;   // cents

// Inputs
let bolts: number = parseInt(prompt("How many bolts would you like to purchase?") || "0");
let nuts: number = parseInt(prompt("How many nuts would you like to purchase?") || "0");
let washers: number = parseInt(prompt("How many washers would you like to purchase?") || "0");

// Output counts
console.log("Number of bolts:          ", bolts);
console.log("Number of nuts:           ", nuts);
console.log("Number of washers:        ", washers);

// Order checks
if (nuts < bolts) {
  console.log("Check the Order - not enough nuts for the bolts you purchased.");
}
if (washers < bolts) {
  console.log("Check the Order - not enough washers for the bolts you purchased.");
}
if (nuts >= bolts && washers >= bolts) {
  console.log("Order is OK.");
}

// Total cost
let totalCost = bolts * BOLT_COST + nuts * NUT_COST + washers * WASHER_COST;
console.log(`Your total cost of the order is ${totalCost} cents.`);

// Done message
console.log("\nDone.");