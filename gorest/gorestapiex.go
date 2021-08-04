package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	College string `json:"college"`
	Number  string `json:"number"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStudents")
	json.NewEncoder(w).Encode(Students)
}

func returnSingleStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Students {
		if article.Id == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func createNewStudent(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Student struct
	// append this to our Student array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student Student
	json.Unmarshal(reqBody, &student)
	// update our global Students array to include
	// our new Student
	Students = append(Students, student)

	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, student := range Students {
		if student.Id == id {
			Students = append(Students[:index], Students[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("students", returnAllStudents)
	myRouter.HandleFunc("/student", createNewStudent).Methods("POST")
	myRouter.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/student/{id}", returnSingleStudent)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	Students = []Student{
		Student{Id: "100", Name: "Saikumar", College: "Jntua", Number: "9866954863"},
		Student{Id: "101", Name: "Manoj", College: "Anna University", Number: "9494109244"},
	}
	handleRequests()
}
