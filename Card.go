package main

import "fmt"

type card struct {
	card string 
} 

func main () {
	var card int
	fmt.Print("=")
	fmt.Scan(&card)
	if card >= 52 {
		fmt.Println(card,"Error")
	}else if card <= 0 {
		fmt.Println(card,"Error")
	}else if card <= 13 {
		fmt.Println(Spades)
	}else if card <= 26 {
		fmt.Println(card,"Diamonds")
	}else if card <= 39 {
		fmt.Println(card,"Clubs")
	}else if card <= 52 {
		fmt.Println("Hearts")
	}
}	

	
