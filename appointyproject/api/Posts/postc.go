package Posts

import (
	"fmt"
	"net/http"
	"sync"
)

type PostAPI struct{}

type Post struct {
	Id       uint64 `json:"id,omitempty"`
	Caption string `json:"username,omitempty"`
	Image    string `json:"email,omitempty"`
	Time string `json:"password,omitempty"`
}

var db []*Post{}
var nextPostId uint64
var lock sync.Mutex

func (p *PostAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
