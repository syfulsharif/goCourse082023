package main

import (
	"log"
	"time"
)

func main() {
	user := User{
		FirstName: "Syful",
		LastName:  "Sharif",
	}

	log.Println(user.FirstName, user.LastName, user.Birthdate, user.Age)
}

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

// func saySomething(s string) (string, string) {
// 	log.Println("s from the saySomething func is ", s)
// 	return s, "world"
// }
