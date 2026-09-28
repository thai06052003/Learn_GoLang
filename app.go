package main

import (
	"bufio"
	"fmt"
	"os"
)

var address = "HCM"
func main() {
	// var fullName = "Dinh Xuan Thai"
	// fmt.Println((fullName))

	// phone := "0852834966"

	// var toan, tiengviet, tunhien int
	// toan =1
	// tiengviet = 2
	// tunhien = 3

	// fmt.Println(phone)
	// fmt.Println(address)
	// fmt.Println(toan)
	// fmt.Println(tiengviet)
	// fmt.Println(tunhien)

	var hoten string

	fmt.Println("Vui long nhap ho ten: ")

	//fmt.Scan(&hoten)

	scanner := bufio.NewScanner(os.Stdin)

	if (scanner.Scan()) {
		hoten =  scanner.Text()
	}

	fmt.Println("Ho ten: ", hoten)
}