package main

import "fmt"

// Membuat variabel
func main() {
	// ----VARIABEL BOOLEAN---//
	var boolean = true
	fmt.Println(boolean)

	boolean = false
	fmt.Println(boolean)
	// ----VARIABEL INT---//
	var numberic = 100
	fmt.Println(numberic)

	numberic = 2022
	fmt.Println(numberic)
	// ----VARIABEL STRING---//
	fullname := "Fuad Azkia"
	fmt.Println(fullname)

	fullname = "Ade"
	fmt.Println(fullname)
	// ----DEKLARASIKAN VARIABEL JAMAK----//
	var (
		firstName = "Michael"
		lastName  = "Jackson"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
