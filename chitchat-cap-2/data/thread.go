package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	Userid    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// Get users that stated the thread

func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", thread.Userid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Format the dates to look nicer on the screen

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in teh chat

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil {
		log.Println("error: not ale to query the number of replies of thread", thread.Id)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			log.Println("error: not able to get the count of the thread")
			return
		}
	}
	return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1", thread.Id)
	if err != nil {
		log.Println("error: not able to query the posts for the thread", thread.Id)
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
		if err != nil {
			log.Println("error: not able to parse the post:", err)
			return
		}
		posts = append(posts, post)
	}
	return
}

func CreateThread(user *User, topic string) (thread Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println("error: prepareing DB statement for create a thread", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.Userid, &thread.CreatedAt)
	if err != nil {
		log.Println("error: add a thread to the db", err)
	}
	return
}

func CreatePost(user *User, thread *Thread, body string) (post Post, err error) {
	statement := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println("error: preparing DB statement for create post", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), body, user.Id, thread.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	if err != nil {
		log.Println("error: insert post to the database", err)
	}
	return
}

// Get all threads from DB

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		log.Println("error: not able to query threads from the table", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		thread := Thread{}
		err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.Userid, &thread.CreatedAt)
		if err != nil {
			log.Println("error: not able to copy thread from db", err)
			return
		}
		threads = append(threads, thread)
	}
	return
}

// Get thread by UUID
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.Userid, &thread.CreatedAt)
	return
}
