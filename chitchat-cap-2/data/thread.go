package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// Return a list of Threads from the database
func Threads() (threads []Thread, err error) {
	// Query the threads on the Data Base
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC ;")
	if err != nil {
		log.Println("error on query threads from database.", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		th := Thread{}
		err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt)
		if err != nil {
			log.Println("error on converting thread from database.", err)
			return
		}
		threads = append(threads, th)
	}
	return
}

// Thread method to return the information abou tthe user
// that created the thread.
func (thread Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1;", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Return the number of posts attached to a thread_id (Number of responses.)
func (thread Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1;", thread.Id)
	if err != nil {
		log.Println("error on query number of replies of thread from database.", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println("error on scaning post count from database.", err)
			return
		}
	}
	return
}
