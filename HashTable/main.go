package main

import (
	htchaining "DSA/HashTable/HT_Chaining"
	ht "DSA/HashTable/HT_Open_Addressing"
)

func main() {

	// Test the hash table implementation
	// HasTable_OpenAddressing()
	// Add more test cases as needed
	HasTable_Chaining()
}

func HasTable_Chaining() {
	// Create a new hash table
	htl := htchaining.CreateTable(10) // Initialize the hash table with a size of 10

	// Insert keys into the hash table
	htl.Insert("apple")
	htl.Insert("banana")
	htl.Insert("orange")
	htl.Insert("grape")
	htl.Insert("grape")

	htl.Insert("grape")

	// Display the hash table
	htl.Display()

	// Search for a key in the hash table
	if htl.Search("grape") {
		println("Key found")
	} else {
		println("Key not found")
	}

	// Delete a key from the hash table
	htl.Delele("grape")
	println("After deletion:")
	htl.Display()
}

func HasTable_OpenAddressing() {
	// Create a new hash table
	htl := ht.CretaeTable(10) // Initialize the hash table with a size of 10

	// Insert keys into the hash table
	htl.Insert("apple")
	htl.Insert("banana")
	htl.Insert("orange")
	htl.Insert("orange")

	// Display the hash table
	htl.Display()

	// Search for a key in the hash table
	index := htl.Search("banana")
	if index != -1 {
		println("Key found at index:", index)
	} else {
		println("Key not found")
	}

	// Delete a key from the hash table
	htl.Delete("banana")
	println("After deletion:")
	htl.Display()
}
