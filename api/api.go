package api

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	arg := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?loc=Local&parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error
	db, err = gorm.Open("mysql", arg)
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
