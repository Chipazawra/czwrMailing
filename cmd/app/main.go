package main

import (
	"net/http"

	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/todo"
)

func main() {
	defer afterStart()
	auth.AddHandlers()
	todo.AddHandlers()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func afterStart() {
	//TODO
}
