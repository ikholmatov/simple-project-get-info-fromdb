package metdohs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	ff "project1/Strucks"
)

func (ff.Customer) Insert(mydb string, info ff.Customer) (error, string) {
	ss, err := sql.Open("postgres", mydb)
	if err != nil {
		panic("Error while opening database from met.go")
	}
	myquery := `Insert into Customers (FirstName,Lastname,Username,Email) values ($1,$2,$3,$4)`
	_, err = ss.Exec(myquery, info.FirstName, info.LastName, info.Username, info.Email)
	if err != nil {
		panic("Error while Inserting values into Customers")
	}

}
func helloe() {
	asd := ff.Customer{FirstName: "davron"}
	fmt.Println(asd)
}
