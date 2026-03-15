package main

import "fmt"

const JwtToken string = "projecthanif"

// can't use := outside of a function
// whoami := "nice"

func main() {
	var username string = "projecthanif"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T\n", isLoggedin)

	var count int = 0
	fmt.Println(count)
	fmt.Printf("Variable is of type: %T\n", count)

	var score float32 = 255.9173738393828292
	fmt.Println(score)
	fmt.Printf("Variable is of type: %T\n", score)

	var score64 float64 = 255.9173738393828292
	fmt.Println(score64)
	fmt.Printf("Variable is of type: %T\n", score64)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type inference
	// the type is inferred from the value assigned
	var wesbite = "https://projecthanif.1hanif.click"
	fmt.Println(wesbite)
	fmt.Printf("Variable is of type: %T\n", wesbite)

	// short declaration no var
	// the type is inferred from the value assigned
	// only works inside functions

	fullName := "Ibrahim Mustapha"
	fmt.Println(fullName)
	fmt.Printf("Variable is of type: %T\n", fullName)

	fmt.Println(JwtToken)
	fmt.Printf("Variable is of type: %T\n", JwtToken)
}
