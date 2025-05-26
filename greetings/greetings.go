package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Greet(name string, age int) (string, int, bool, error) {

	var ageMessage string
	var isAdult bool
	

	if age < 18 {
		ageMessage  = "Im sorry, you cant enter this site cause you're still young"
		isAdult  = false
	} else if age >= 18 && age <= 30 {
		ageMessage = "Welcome to our adult website where you can learn how to code with hot woman"
		isAdult = true
	} else {
		ageMessage = "Im sorry but you're too old for this"
	}

	if name == "" {
		return "", 0, false, errors.New("empty name")
	}

	message := fmt.Sprintf(randomMessage(), name, ageMessage)
	
	return message, age, isAdult, nil

}

func randomMessage() string {
	formats := []string{
		"Hi, %s, %s",
		"Hello, %s %s",
		"hallow, %s %s",
	}

	return formats[rand.Intn(len(formats))]
}