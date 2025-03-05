package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	fmt.Println("got /users")
	io.WriteString(response, "Este es un endpoint de users")
	fmt.Fprintln(response, " Endpoint GetUsers")
}

func getCourses(response http.ResponseWriter, request *http.Request) {
	fmt.Println("got /users")
	io.WriteString(response, "Endpoint de courses")
	fmt.Fprintln(response, "Endpoint de coursos")
}
