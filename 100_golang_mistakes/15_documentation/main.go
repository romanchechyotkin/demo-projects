package main

import (
	"log"

	"mistakes/15_documentation/customer"
)

// DefaultPermission is the default permission used by the store engine.
const DefaultPermission = 0o644 // Need read and write accesses.

func main() {
	c := customer.NewCustomer("1")
	c2 := customer.New("2")
	log.Println(c.ID())
	log.Println(c2.ID())
}
