//generating a random number
package main

import (
	"fmt"
	"math/rand"
	"time"

	
)

func main() {

	seconds := time.Now().Unix() // it generates random numbers at different time zone
	rand.Seed(seconds)
	random := rand.Intn(100) +1 // generate any number from 1- 100
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("can you guess it?")
	var number int


	var success bool
	
	for num := 0; num < 10;num++{
	fmt.Println("you have", 10-num, "guesses left. ")
	
	fmt.Println("Make a guess: ")
	_, err := fmt.Scanln(&number)
	if err != nil{
		panic(err)
	}
	if number < random {
		fmt.Println("Oops. Your guess is LOW.")
	}else if number > random{
		fmt.Println("Oops. Your guess is HIGH ")
	}else{
		success = true
		fmt.Println("Good job! You guessed it!")
		break
	}
	
}

if !success{
	fmt.Println("sorry, you didnt guess my number. It was", random)
}

	
	

}