// Aquí estamos limitando el tipo de usuario a las const que ya definimos en el enum.

package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {
	eduardo := User{Username: "Edu", Type: Teacher}
	uriel := User{Username: "Uri", Type: Student}
	fmt.Println(eduardo)
	fmt.Println(uriel)
}
