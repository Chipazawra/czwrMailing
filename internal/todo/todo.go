package todo

import (
	"fmt"
	"net/http"
)

func AddHandlers() {
	http.HandleFunc("/todo", Pong)
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "todo")
}
