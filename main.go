package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"hash/fnv"
	"log"
	"time"

)

func (Customer) Insert(mybase string, a Customer) (error, string) {
	db, err := sql.Open("postgres", mybase)
	if err != nil {
		log.Panicf("%s\n%s", "Error While opening DB", err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil{
		log.Panicf("%s\n%s", "Error While opening Transacion", err)
	}

	uuidCustomer := uuid.New()
	QueryCustomers := `Insert into customers Values($1,$2,%3,%4,$5,$6,$7,$8,$9)`
	_,err = tx.Exec(QueryCustomers, uuidCustomer.String(), a.FirstName, a.LastName, a.UserName, a.Email, a.Gender, a.BirthDate, a.Password, a.Status)
	if err != nil{
		tx.Rollback()
		return
	}
	for _,v := range a.Phones{
		uuidPhone := uuid.New()
		QueryPhone := `Insert into phones values($1,$2,$3,$4)`
		_,err := tx.Exec(QueryPhone,uuidPhone,uuidCustomer,pq.Array(v.Numbers),v.Code)
		if err != nil {
			tx.Rollback()
			return
		}
	}






	return err, "ok"
}
func main() {
	object := Customer{
		FirstName: "Davron",
		LastName:  "Kholmatov",
		UserName:  "venomuz",
		Phones: []Phone{
			{Numbers: []int64{909331343, 977535979}, Code: "+990"},
			{Numbers: []int64{123456755, 852852885}, Code: "+1"},
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
		Email:     "venom.uzz@mail.ru",
		Gender:    "M",
		BirthDate: time.Date(2002, time.May, 26, 00, 00, 00, 00, time.UTC),
		Password: func() uint32 {
				pass := "root12345"
				h := fnv.New32a()
				h.Write([]byte(pass))
				return h.Sum32()
			}(),
		Status: true,
		}
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
	Password  uint32 // should be hashed and validate password should be 8 symbols
	Status    bool
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
