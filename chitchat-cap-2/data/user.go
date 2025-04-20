package data

import (
	"log"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user User) IsEmpty() bool {
	return user.Uuid == ""
}

func UserByEmail(email string) (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1;", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func CreateUser(name string, email string, password string) (user User) {
	user = User{}
	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, name, email, password, created_at;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println("error creating user", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), name, email, Encrypt(password), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Println("error getting the users information upon creation.", err)
	}
	return
}
