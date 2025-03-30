package main

import (
	"fmt"
	"net/http"
)

func main() {
	HandleRequest()
}
func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactsPage)
	http.ListenAndServe(":8080", nil)
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func ContactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts Page")
}
