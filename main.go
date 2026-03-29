package main

import (
	"fmt"
	
)

func ConvertKilogramtoGram(n float64) float64{
		return n*1000

}

func ConvertGramToKilogram(num float64) float64{
	return num/1000
}

func main() {
	var choice int
	var in, swi float64
	

	fmt.Println("=====Modern Weight Converter====")
	fmt.Printf("1...continue......2...Exit\n")


		fmt.Scanln(&choice)

		switch choice {
		case 1:
	
		case 2:
			return
		default:
		if choice > 2{
			fmt.Println("option 1 or 2 only available")
			return
			}
		}
	

	for {

		fmt.Println()
		fmt.Println("Welcome Buddy!!")
		fmt.Printf("Select An Option\n 1...Convert Kilogram To Gram\n 2...Convert Gram To Kilogram\n 3...Exit\n")
		fmt.Scanln(&in)
		
		switch in {

		case 1:
			fmt.Println()
			fmt.Println("===Please Enter Digit To Be Converted To Gram===")
			fmt.Scanln(&swi)
			fmt.Println(ConvertKilogramtoGram(swi),"Gram")
			fmt.Println()
			break

		case 2:
			fmt.Println()
			fmt.Println("===Please Enter A Digit To Be Converted To Kilogram")
			fmt.Scanln(&swi)
			fmt.Println(ConvertGramToKilogram(swi),"Kilogram")
			continue

		case 3:
			return
		}
	}

}