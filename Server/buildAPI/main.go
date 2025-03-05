package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//model for course --file
type Course struct {
	CourseId string `json:"course_id"`
	CourseName     string `json:"course_name"`
	Price    float64 `json:"price"`
	Author *Author `json:"author"`

}

type Author struct {
	FullName string `json:"full_name"`
	Website string `json:"website"`
}

//fake db
var Courses []Course

//middleware / helper functions

func (c *Course) isEmpty() bool{
	return c.CourseId == "" && c.CourseName == ""

}

func main() {
	fmt.Println("Hello, World!")
}

// Controllers --file
func serverHome (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to the server</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Courses)
}

func getOneCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range Courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty(){
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	// generate unique id (string)
	// append value to the slice 

	course.CourseId = strconv.Itoa(rand.Intn(1000000))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range Courses{
		if course.CourseId == params["id"]{
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range Courses{
		if course.CourseId == params["id"]{
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return
}

