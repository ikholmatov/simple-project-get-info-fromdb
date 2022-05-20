package metdohs

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/venomuz/project1/models"
)

func (a models.Customer) Insert(mydb string, info models.Customer) (error, string) {
	ss, err := sql.Open("postgres", mydb)
	if err != nil {
		panic("Error while opening database from met.go")
	}
	myquery := `Insert into Customers (FirstName,Lastname,Username,Email,Gender,Birthdate,Password,Status) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err = ss.Exec(myquery, info.FirstName, info.LastName, info.Username, info.Email, info.Gender, info.Status)
	if err != nil {
		panic("Error while Inserting values into Customers")
	}

	return err, ""
}
