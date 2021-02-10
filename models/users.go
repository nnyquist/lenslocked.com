package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewUserService connects to the database
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil
}

// UserService allows us to interact with User objects
type UserService struct {
	db *gorm.DB
}

// ErrNotFound is returned when a resource cannot be found in the database
var ErrNotFound = errors.New("models: resource not found")

// ByID will look up a user by the id provided.
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError --> something went wrong server side or in code
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Close the UserService database connection
func (us *UserService) Close() error {
	sqlDB, _ := us.db.DB()
	return sqlDB.Close()
}

// DestructiveReset drops the users table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.Migrator().DropTable(&User{})
	us.db.AutoMigrate(&User{})
}

// User is ...
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}
