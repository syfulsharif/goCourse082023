package main

import "fmt"

func main() {
	fmt.Println("Hello Go World")

	var whatToSay string
	var i int32

	whatToSay = "Goodbye cruel world"
	i = 50
	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"

}
