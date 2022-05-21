package repo


import(
	"github.com/venomuz/project1/models"
)

type Mtehodss interface{
	Insert(info *models.Customer) (error, string)
}