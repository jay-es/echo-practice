package api

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type User struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Birthdate time.Time `json:"birthdate" gorm:"type:date"`
	// Added      time.Time `gorm:"type:timestamp"`
}

type Person struct {
	User
}

func init() {
	// parseTime=True
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost)/test?parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
}

func GetUsers() []User {
	var users []User
	db.Find(&users) // SELECT * FROM authors;
	return users
}

func GetUserById(id string) User {
	var user User
	db.First(&user, id) // SELECT * FROM authors WHERE id = 10;
	return user
}

func GetUserByName(name string) User {
	var user User
	db.Where("first_name = ?", name).First(&user)
	return user
}

// other name struct
func GetPerson() Person {
	var person Person
	db.Table("users").First(&person)
	return person
}
