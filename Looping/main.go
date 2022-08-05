package main

import "log"

func main() {
	//animals := []string{"dog","fish","horse","cow","cat"}
	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"
	// var firstLine = "Once upon a midnight dreary"
	// firstLine = "x"

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}
	
	var users []User

	users = append(users, User{"John","Smith","john@smith.com",30})
	users = append(users, User{"Marry","Jones","marry@smith.com",20})
	users = append(users, User{"Sally","Brown","sally@smith.com",45})
	users = append(users, User{"Alex","Anderson","alex@smith.com",18})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}