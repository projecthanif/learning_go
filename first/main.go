package main

import "fmt"

func main() {
	const age int = 23
	var firstName string = "Ibrahim"
	var lastName string = "Mustapha"
	var username string = "projecthanif"
	var githubUrl string = "https://github.com/" + username

	fmt.Println("Hello World!!")
	fmt.Printf("My name is %s %s and I am %d years old.\nMy github username is %s\n", firstName, lastName, age, username)
	fmt.Printf("My github URL is %s\n", githubUrl)
}
