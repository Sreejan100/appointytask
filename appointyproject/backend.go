package main

import (
	"log"
	"net/http"
	"/home/sreejan/Documents/appointyproject/api/users"
	"/home/sreejan/Documents/appointyproject/api/Posts"

)

func main() {
	http.HandleFunc("/users/", &user.UserAPI{})
	http.HandleFunc("/users/post", &postc.PostAPI{})
	log.Fatal(http.ListenAndServe(":27017", nil))
}
