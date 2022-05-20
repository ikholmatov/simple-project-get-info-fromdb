package models

import "time"

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Phones    []Phone
	Adresses  []Adress
	Products  []Product
	Email     string
	Gender    string
	Birthdate time.Time
	Password  string // should be hashed and validate password should be 8 symbols
	Status    string
}

type Phone struct {
	ID      string
	UserID  string
	Numbers []int64 // 998812891, 998802891 ...
	Code    string  // "+998"
}

type Adress struct {
	ID          string
	Country     string
	City        string
	District    string
	PostalCodes []int64
}

type Product struct {
	ID          string
	Name        string
	Cost        int64
	OrderNumber int64
	Amount      int64
	Currency    string
	Rating      int64
}
