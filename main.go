package main

import (
	"fmt"
	"go-backend/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Charles","David","Casenas"}

	messages, error := greetings.GreetToMultipleUser(names)

	// message, _, _, error := greetings.Greet("cc", 56)
	// fmt.Println(message)
	
	if error != nil {
		fmt.Println("Error:", error)
	}

	fmt.Println(messages)
}