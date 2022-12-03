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


	var success bool // a sucess variable 
	
	for num := 0; num < 10;num++{ // our users to guess up to 10 times
	fmt.Println("you have", 10-num, "guesses left. ")
	
	fmt.Println("Make a guess: ")
	_, err := fmt.Scanln(&number) // accept an input from the command line interface
	if err != nil{
		panic(err)
	}
	if number < random {// checks if the guess number is lesss than than the random number
		fmt.Println("Oops. Your guess is LOW.")
	}else if number > random{ // checks if the guess number is greater than than the random number
		fmt.Println("Oops. Your guess is HIGH ")
	}else{
		success = true // success is true then user guessed right
		fmt.Println("Good job! You guessed it!")
		break
	}
	
}

if !success{ // success is false then user guessed wrong
	fmt.Println("sorry, you didnt guess my number. It was", random)
}

	
	

}