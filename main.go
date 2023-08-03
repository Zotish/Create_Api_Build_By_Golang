package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	Name     string `json:"name"`
	CourseId int    `json:"id"`
}

var courses []Course

func main() {
	fmt.Println("hello")
	r := mux.NewRouter()
	courses = append(courses, Course{Name: "Zotish", CourseId: 3})
	courses = append(courses, Course{Name: "Jannat", CourseId: 2})
	r.HandleFunc("/", Servehome).Methods("GET")
	r.HandleFunc("/products", GetAllCourse).Methods("GET")
	r.HandleFunc("/product/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/product", CreatOneCourse).Methods("POST")
	r.HandleFunc("/product/{id}", UpdateoneCourse).Methods("PUT")
	r.HandleFunc("/product/{id}", Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9002", r))
}
func Servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to our"))
}
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	iSTr := params["id"]
	convT, _ := strconv.Atoi(iSTr)
	for _, course := range courses {
		if course.CourseId == convT {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}
func CreatOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data available ")
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.Name == "" {
		json.NewEncoder(w).Encode("Send some data please")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = rand.Intn(100)
	json.NewEncoder(w).Encode(course)

}
func UpdateoneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	iStr := params["id"]
	strc, _ := strconv.Atoi(iStr)
	for key, course := range courses {
		if course.CourseId == strc {
			courses = append(courses[:key], courses[key+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = strc
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Id is not found ")
}
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	iStr := params["id"]
	IntT, _ := strconv.Atoi(iStr)
	for key, course := range courses {
		if course.CourseId == IntT {
			courses = append(courses[:key], courses[key+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode("No id found To delete")
}
func DeeletAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//Delete all body value is easy
}
