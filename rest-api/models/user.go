package models

import (
	"github.com/alopes2/go-studies/rest-api/db"
	"github.com/alopes2/go-studies/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {

	query := `
	INSERT INTO users(email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassowrd(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	user.ID = userID

	return err
}
