package model

import (
	"database/sql"
	"log"
	"web/db"
)

var databaseconn  *sql.DB  = db.DbConnection()

type User struct {
	Id int64
	Username string `validate:"required"`
	Email string `validate:"required,email"`
	Password string `validate:"required"`
}

func (u User) Save(){ 
	stmt, err := databaseconn.Prepare("INSERT INTO users(username, email, password_hash) VALUES(?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(u.Username, u.Email, u.Password)
}

func (u User) Update(){ 
	stmt, err := databaseconn.Prepare("UPDATE users SET username = ?, email = ?, password_hash = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	log.Printf("Executing query: UPDATE users SET username = '%s', email = '%s', password_hash = '%s' WHERE id = %d",
	u.Username, u.Email, u.Password, u.Id)
	stmt.Exec(u.Username, u.Email, u.Password, u.Id)
}

func (u User) Delete(){ 
	stmt, err := databaseconn.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec(u.Id)
}

func FindUserByID(id int64) User {
	return User {Id: 2}
}

