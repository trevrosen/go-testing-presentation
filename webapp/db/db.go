package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID        string `db:id`
	FirstName string
	LastName  string
	Age       int
}

type CommonDBInterface interface {
	FindUserById(id string) (User, error)
}

// RealDB implements CommonDBInterface
type RealDB struct {
	Gorm *gorm.DB
}

// NewDB returns an intialized database
func NewDB() (RealDB, error) {
	var err error
	newDB := RealDB{}

	newDB.Gorm, err = gorm.Open("sqlite3", "demo.sqlite3")
	newDB.Gorm.AutoMigrate(&User{})
	return newDB, err
}

func (rdb RealDB) FindUserById(id string) (User, error) {
	user := User{}
	var returnableErr error
	if err := rdb.Gorm.Where("id = ?", id).First(&user); err != nil {
		returnableErr = fmt.Errorf("User not found")
	}
	return user, returnableErr
}

// MockDB implements CommonDBInterface
type MockDB struct {
	OnFindUserById func(string) (User, error)
}

func (mdb MockDB) FindUserById(id string) (User, error) {
	return mdb.OnFindUserById(id)
}
