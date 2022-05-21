package main

import (
	"database/sql"
	"fmt"
	m "github.com/venomuz/project1/metdohs"
	s "github.com/venomuz/project1/repo"
	md "github.com/venomuz/project1/models"
)

type Mtds struct{
	storage *s.Mtehodss
}

func main() {

	mybase := "user=kilogram password=112233 dbname=kilogram sslmode=disable"
	db, err := sql.Open("postgres", mybase)
	if err != nil {
		fmt.Println(err)
		return
	}
	mr  := m.NewReviewRepo(db)
	ff := md.Customer{Username: "asd"}


	mr.Insert(&ff)
	
}
