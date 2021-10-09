package users

import (
	"fmt"
	"net/http"
	"sync"
)

type UserAPI struct{}

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var db []*User{}
var nextUserId uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method '%v' to '%v'\n", r.Method, r.URL)
	}
}
