package main

import (
	"fmt"
	"os"
)
	
var letter = [5] string {"M", "S", "A", "D", "E"}

func main() {
	exit := false
	for exit == false {

		fmt.Println("(M) Multiplication", "(S) Subtract", "(A) Addition", "(D) Division", "(E) Exit", "\n Press the Button what you want to do!")

		var choice string
		fmt.Scan(&choice)

		if choice == letter[4] {
			fmt.Println("Thank you, have a nice Day!")
			exit = true
			os.Exit(0)
		}

		fmt.Println("Please enter the first Number ->")

		var numOne int
		fmt.Scanln(&numOne)
		fmt.Println("Please enter the second Number ->")
	
		var numTwo int
		fmt.Scanln(&numTwo)

		if choice == letter[0] {
		fmt.Println("Answer:", numOne * numTwo)
		}

		if choice == letter[1] {
			fmt.Println("Answer:", numOne - numTwo)
		}

		if choice == letter[2] {
			fmt.Println("Answer:", numOne + numTwo)
		}

		if choice == letter[3] {
			fmt.Println("Answer:", numOne / numTwo)
		}

		
	}
}	