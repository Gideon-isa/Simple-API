// Package main implements the basic API

// A simple Create, Read, Update and Delete API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is the Author full details
type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is the body of the response
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// A slice of Post struct
var postList []Post = []Post{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	fmt.Println("Starting the server...")
	err := http.ListenAndServe(":8000", router)

	if err != nil {
		log.Fatal("Error opening the server")
	}

}

// addItem adds Post object to the post slice
func addItem(w http.ResponseWriter, r *http.Request) {
	// get the Item Value from the JSON body
	w.Header().Set("Content-Type", "json")
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	postList = append(postList, newPost)
	json.NewEncoder(w).Encode(postList)

}

// getAllPost returns the slice of posts
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	json.NewEncoder(w).Encode(postList)
}

// getsPost returns the specified Post
func getPost(w http.ResponseWriter, r *http.Request) {
	//get the ID of the post from the root parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID could not be converted to an integer"))
		return
	}

	if id >= len(postList) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	postID := postList[id]

	w.Header().Set("content-Type", "json")
	json.NewEncoder(w).Encode(postID)

}

func updatePost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(postList) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// get the value from the JSON body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	postList[id] = updatedPost

	w.Header().Set("Content-Type", "json")
	json.NewEncoder(w).Encode(updatedPost)

}

// patchPost update a Post
func patchPost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(postList) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// get the current value
	post := &postList[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "json")
	json.NewEncoder(w).Encode(post)
}

// deletePost removes the specified post
func deletePost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(postList) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	// Delete the post from the slice
	postList = append(postList[:id], postList[id+1:]...)

	w.WriteHeader(200)
}
