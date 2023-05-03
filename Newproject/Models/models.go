package models
import(
	"gorm.io/gorm"
)
type Employee struct{
	gorm.models
	id string  'json:"id"'
	name string 'json:"name"'
	salary int 'json:"salary"'
	project *Project 'json:"id"'
	maneger *Maneger 'json:"maneger"'
	}
	type Product struct {
		gorm.Model
		Name  string `json:"name" gorm:"text;not null;default:null`
		Price string `json:"price" gorm:"text;not null;default:null`
	}
	

	type maneger struct{
		gorm.models
		id string 'json:"id"'
		name string 'json:"name"'
	}


