package metdohs

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/venomuz/project1/models"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{
		db: db,
	}
}

func (r *ReviewRepo) Insert(info *models.Customer) (error, string) {
	myquery := `Insert into Customers (FirstName,Lastname,Username,Email,Gender,Birthdate,Password,Status) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := r.db.Exec(myquery, info.FirstName, info.LastName, info.Username, info.Email, info.Gender, info.Status)
	if err != nil {
		panic("Error while Inserting values into Customers")
	}

	return err, ""
}
