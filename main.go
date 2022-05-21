package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func (Customer) Insert(mybase string, UserDate Customer) (error, string) {
	db, err := sql.Open("postgres", mybase)
	if err != nil {
		log.Panicf("%s\n%s", "Error While opening DB", err)
	}
	defer db.Close()

	QueryCustomers := `Insert into customers Values($1,$2,%3,%4,$5,$6,$7,$8,$9)`
	return err, "ok"
}
func main() {
	object := Customer{
		FirstName: "Davron",
		LastName:  "Kholmatov",
		UserName:  "venomuz",
		Phones: []Phone{
			{Numbers: []int64{909331343, 977535979}, Code: "+990"},
		},
		Addresses: []Address{{
			Country:     "Uzbekistan",
			City:        "Tashkent",
			District:    "Uch Tepa",
			PostalCodes: []int64{100123, 100193},
		},
		},
		Products: []Product{{
			Name: "T-Shirt",
			Types: []Type{
				{
					Name: "Dress",
				},
				{
					Name: "Clothing",
				},
			},
			Cost:        99000,
			OrderNumber: 1,
			Amount:      3,
			Currency:    "Dollar",
			Rating:      5,
		},
		},
	}
	mybase := "user=kilogram password=112233 dbname=project1 sslmode=disable"

	err, s := object.Insert(mybase, object)
	if err != nil {
		log.Panicf("%s\n%s", "Error While useing Insert method DB", err)
	}
	fmt.Println(s)
}

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Phones    []Phone
	Addresses []Address
	Products  []Product
	Email     string
	Gender    string
	BirthDate time.Time
	Password  string // should be hashed and validate password should be 8 symbols
	Status    string
}

type Phone struct {
	ID         string
	CustomerID string
	Numbers    []int64 // 998812891, 998802891 ...
	Code       string  // "+998"
}

type Address struct {
	ID          string
	CustomerID  string
	Country     string
	City        string
	District    string
	PostalCodes []int64
}

type Product struct {
	ID          string
	CustomerID  string
	Name        string
	Types       []Type
	Cost        int64
	OrderNumber int64
	Amount      int64
	Currency    string
	Rating      int64
}

type Type struct {
	ID        string
	ProductID string
	Name      string
}
