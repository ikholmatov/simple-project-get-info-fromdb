package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"hash/fnv"
	"log"
	"time"
)

func (Customer) Insert(MyBase string, a Customer) (error, string) {
	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		log.Panicf("%s \n %s", "Error While opening DB", err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Panicf("%s \n %s", "Error While opening Transacion", err)
	}

	uuidCustomer := uuid.New()
	QueryCustomers := `Insert into customers Values($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err = tx.Exec(QueryCustomers, uuidCustomer.String(), a.FirstName, a.LastName, a.UserName, a.Email, a.Gender, a.BirthDate, a.Password, a.Status)
	if err != nil {
		tx.Rollback()
		log.Panicf("%s \n %s", "Error While inserting colums to the customers", err)
	}
	for _, v := range a.Phones {
		uuidPhone := uuid.New()
		QueryPhone := `Insert into phones values($1,$2,$3,$4)`
		_, err := tx.Exec(QueryPhone, uuidPhone, uuidCustomer, pq.Array(v.Numbers), v.Code)
		if err != nil {
			tx.Rollback()
			log.Panicf("%s \n %s", "Error While inserting colums to the Phones", err)
		}
	}
	for _, v := range a.Addresses {
		uuidAddress := uuid.New()
		QueryAddress := "Insert into addresses values($1,$2,$3,$4,$5,$6)"
		_, err := tx.Exec(QueryAddress, uuidAddress, uuidCustomer, v.Country, v.City, v.District, pq.Array(v.PostalCodes))
		if err != nil {
			tx.Rollback()
			log.Panicf("%s \n %s", "Error While inserting colums to the customers", err)
		}
	}
	for _, v := range a.Products {
		uuidProduct := uuid.New()
		QueryProduct := "Insert into products values($1,$2,$3,$4,$5,$6,$7,$8)"
		_, err := tx.Exec(QueryProduct, uuidProduct, uuidCustomer, v.Name, v.Cost, v.OrderNumber, v.Amount, v.Currency, v.Rating)
		if err != nil {
			tx.Rollback()
			log.Panicf("%s \n %s", "Error While inserting colums to the products", err)
		}
		uuidType := uuid.New()
		var asd []string
		QueryType := `Insert into types (ID,ProductId,Name) values($1,$2,$3)`
		asd = append(asd, v.Types[0].Name)
		_, err = tx.Exec(QueryType, uuidType, uuidProduct, pq.Array(asd))
		if err != nil {
			tx.Rollback()
			log.Panicf("%s \n %s", "Error While inserting colums to the types", err)
		}

	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return err, "ok"
}
func (Customer) Get(MyBase string, TarID string) {
	var (
		customer Customer
		adress   []Address
		product  []Product
		types    []Type
	)

	db, err := sql.Open("postgres", MyBase)
	if err != nil {
		log.Panicf("%s \n %s", "Error While opening DB", err)
	}
	defer db.Close()
	rows, err := db.Query(`Select * from Customers where CustomerID = $1 ;`, TarID)
	if err != nil {
		log.Panicf("%s \n %s", "Error While get colums from Customers", err)
	}
	for rows.Next() {
		err = rows.Scan(&f.ID, &f.FirstName, &f.LastName, &f.UserName, &f.Email, &f.Gender, &f.BirthDate, &f.Password, &f.Status)
		if err != nil {
			log.Panicf("%s \n %s", "Error While Scaning colums from Customers", err)
		}
	}
	defer rows.Close()

	for rows.Next() {
		rows, err = db.Query(`Select * from Phones where CustomerID = $1 ;`, TarID)
		if err != nil {
			log.Panicf("%s \n %s", "Error While get colums from Phones", err)
		}
		err = rows.Scan(&num.ID, &num.CustomerID, &num.Numbers, &num.Code)
		if err != nil {
			log.Panicf("%s \n %s", "Error While Scaning colums from Phones", err)
		}
	}

	for rows.Next() {
		rows, err = db.Query(`Select * from addresses where CustomerID = $1 ;`, TarID)
		if err != nil {
			log.Panicf("%s \n %s", "Error While get colums from addresses", err)
		}
		err = rows.Scan(&adrs.ID, &adrs.CustomerID, &adrs.Country, &adrs.City, &adrs.District, &adrs.PostalCodes)
		if err != nil {
			log.Panicf("%s \n %s", "Error While Scaning colums from addresses", err)
		}
	}

	for rows.Next() {
		rows, err = db.Query(`Select * from products where CustomerID = $1 ;`, TarID)
		if err != nil {
			log.Panicf("%s \n %s", "Error While get colums from products", err)
		}
		err = rows.Scan(&prod.ID, &prod.CustomerID, &prod.Name, &prod.Cost, &prod.OrderNumber, &prod.Amount, &prod.Currency, &prod.Rating)
		if err != nil {
			log.Panicf("%s \n %s", "Error While Scaning colums from products", err)
		}
	}

	for rows.Next() {
		rows, err = db.Query(`Select * from types where CustomerID = $1 ;`, TarID)
		if err != nil {
			log.Panicf("%s \n %s", "Error While get colums from types", err)
		}
		err = rows.Scan(&typ.ID, &prod.CustomerID, &typ.Name, &typ.Name)
		if err != nil {
			log.Panicf("%s \n %s", "Error While Scaning colums from types", err)
		}
	}
	//f.Phones = append(f.Phones, num)
	//f.Addresses = append(f.Addresses, adrs)
	//prod.Types = append(prod.Types, typ)
	//f.Products = append(f.Products, prod)
	ans, err := json.MarshalIndent(f, "", "   ")
	if err != nil {
		log.Panicf("%s \n %s", "Error While Marshaling colums", err)
	}
	fmt.Println(string(ans))
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
	mybase := "user=venom password=112233 dbname=project1 sslmode=disable"
	object.Get(mybase, "74c739a6-372e-441e-8cb8-2fde9d01c574")
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
