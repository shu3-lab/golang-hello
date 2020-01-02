package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", auth)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func auth(w http.ResponseWriter, r *http.Request) {
	var (
		user string
		pass string
	)
	user = "user"
	pass = "pass"
	if !checkAuth(r, user, pass) {
		w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Not authorized", 401)
		return
	}
	w.Write([]byte("Authentification is success!"))

}

func checkAuth(r *http.Request, u string, p string) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return username == u && password == p
}
