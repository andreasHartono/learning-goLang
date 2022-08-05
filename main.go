package main

import (
	"log"
	"time"
	"fmt"

)

//import "fmt"
var s = "seven"
var firstName string
var lastName string
var age int
var birthDate time.Time

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time	
}


func main() {
	user := User {
		FirstName: "Andreas",
		LastName: "Hartono",
		PhoneNumber: "1 555- 555-1212",
	}
	fmt.Println(user.PhoneNumber)
	log.Println(user.FirstName, user.LastName, "Birthdate:",user.BirthDate)
}