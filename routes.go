package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	testPost := Post{
		Id: 1, Title: "Title 1", Text: "Text 1",
	}
	posts = []Post{testPost}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling data"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)

}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var post Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		res.Write([]byte(`{"error": "Error unmarshalling request data"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling data"}`))
		fmt.Println(err)
		return
	}
	res.Write(result)
}
