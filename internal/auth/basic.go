package auth

import (
	"fmt"
	"net/http"
)

func AddHandlers() {
	http.HandleFunc("/login", BasicAuth)
	http.HandleFunc("/logout", Logout)
}

func BasicAuth(w http.ResponseWriter, r *http.Request) {

	if usr, pwd, ok := r.BasicAuth(); !ok {

		fmt.Fprint(w, "Error parsing basic auth")
		w.WriteHeader(401)
		return

	} else if usr != "14" || pwd != "88" {

		fmt.Fprint(w, "Wrong login or password...")
		w.WriteHeader(401)
		return

	} else {

		fmt.Fprintf(w, "Hello, %v!", usr)
		w.WriteHeader(401)
		return

	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout")
}
