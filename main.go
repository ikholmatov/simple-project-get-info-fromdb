package main

import (
	"fmt"
	ss "project1/Strucks"
)

func main() {
	person := ss.Customer{FirstName: "Davron", LastName: "Kholmatov", Username: "venom", Phones: []ss.Phone{{UserID: "asd"}}}
	fmt.Println(person)
}
