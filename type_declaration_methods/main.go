package main

import "fmt"

type email string

func (e email) print() {
	fmt.Println("email:", e)
}

type abcEmail email

func main() {
	john := abcEmail("JohnDoe@abc.com")

	// john.print() // john.print undefined (type abcEmail has no field or method print)

	email(john).print() // email: JohnDoe@abc.com

}
