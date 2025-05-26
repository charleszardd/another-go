package main

import (
	"fmt"
	"go-backend/greetings"
)

func main() {

	message, _, _, error := greetings.Greet("cc", 56)
	fmt.Println(message)
	
	if error != nil {
		fmt.Println("Error:", error)
	}
}