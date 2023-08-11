package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int64

	whatToSay = "Goodbye cruel world"
	fmt.Println(whatToSay)

	i = 35
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
