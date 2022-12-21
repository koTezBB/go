package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. he is %d", u.name, u.age)

}

func (u *User) setNewName(newName string) {
	u.name = newName

}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"bob", 25, -45, 4.5, 0.6}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}

func handleRequests() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":8080", nil)
}

func main() {
	//bob := User{name: "bob", age: 25, money: -45, avg_grades: 4.5, happiness: 0.6}

	handleRequests()
}
