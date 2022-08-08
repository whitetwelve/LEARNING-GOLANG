package main

import (
	"fmt"
	"math"
)

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
	// // ----VARIABEL STRING---//
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
	// // CONST JAMAK //
	const (
		fullName string = "Fuad Azkia"
		age      int    = 30
		address  string = "Bekasi"
	)

	fmt.Println(fullName)
	fmt.Println(age)
	fmt.Println(address)

	//--OPERATOR--//
	var result = 10 + 100
	fmt.Println(result)

	const (
		a int = 200
		b int = 10
		c     = a * b
	)
	fmt.Println(c)
	// --PEMBULATAN--//
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Round(1.4))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
	// PENGKONDISIAN VALUE BOOL //
	var name1 = "Fuad"
	var name2 = "Fuad"

	var hasil bool = name1 == name2
	fmt.Println(hasil)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

	// LOOPING //
	for i := 0; i <= 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
}
