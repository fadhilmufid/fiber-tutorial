package user

import (
	"errors"

	"github.com/fadhilmufid/fiber-tutorial/database"
)


func FindById(id int) (User, error) {
	var user User

	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return user, errors.New("user does not exist")
	}
	return user, nil
}

func FindAll() *[]User{
	users := []User{}

	database.Database.Db.Find(&users)
	return &users
}

func CreateData(user *User) {
	database.Database.Db.Create(&user)
}

