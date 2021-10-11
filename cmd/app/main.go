package main

import (
	"net/http"

	"github.com/Chipazawra/czwrmailing/internal/auth"
	"github.com/Chipazawra/czwrmailing/internal/todo"
)

func main() {
	auth.AddHandlers()
	todo.AddHandlers()
	http.ListenAndServe(":8080", nil)
}
