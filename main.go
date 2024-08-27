package main

import (
	"fmt"
)

var balance float64 = 1000

func main() {
	ch := 1
	for ch == 1 {
		fmt.Print("ENTER THE OPERATION YOU WANT TO PERFORM \n 1 : CHECK \n 2 : WITHDRAW \n 3 : DESPOSIT \n")
		var c int
		fmt.Scanln(&c)
		switch c {
		case 1:
			check()
			break
		case 2:
			withdraw()
			break
		case 3:
			deposit()
			break
		}
		fmt.Print("\nDO YOU WANT TO OPERATE AGAIN \n")
		fmt.Scanln(&ch)
	}
}
func check() {
	fmt.Print(balance)
}
func withdraw() {
	var w float64
	fmt.Print("ENTER AMOUNT TO WITHDRAW : ")
	fmt.Scanln(&w)
	if balance >= w {
		balance = balance - w
	} else {
		fmt.Print("INSUFFICIENT BALANCE")
	}
}
func deposit() {
	var d float64
	fmt.Print("ENTER AMOUNT TO DEPOSIT : ")
	fmt.Scanln(&d)
	balance = balance + d
}
