package models

import (
	"encoding/json"
	d "github.com/connyay/goblog/db"
)

type Post struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

var db = d.GetDB()

func GetPost(post_id string) []byte {
	var st = d.GetPrepared("Post", "GetPost", "SELECT id, body FROM posts WHERE id = ?")

	var body string
	var id int
	st.QueryRow(post_id).Scan(&id, &body)
	post := Post{Id: id, Body: body}

	json_post, _ := json.Marshal(post)
	return json_post
}

func GetPosts() []byte {
	var posts []Post
	var st = d.GetPrepared("Post", "GetPosts", "SELECT id, body FROM posts")

	rows, err := st.Query()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var body string
		var id int
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		posts = append(posts, Post{Id: id, Body: body})
	}
	json_posts, _ := json.Marshal(posts)
	return json_posts
}
