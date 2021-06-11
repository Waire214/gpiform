package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Details struct {
	Email    string
	Password string
}

var allData []Details
var tpl *template.Template

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "postform.html", nil)
}
func processPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processPostHandler is running")
	fmt.Println(allData)
	var formData Details
	formData.Email = r.FormValue("useremail")
	formData.Password = r.FormValue("userpassword")
	allData = append(allData, formData)
	fmt.Println("Email: ", formData.Email, "Password: ", formData.Password)
	fmt.Println(allData)
	json.NewEncoder(w).Encode(formData)
	json.NewEncoder(w).Encode(allData)
}
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/postform", postFormHandler)
	http.HandleFunc("/processpost", processPostHandler)
	log.Fatal(http.ListenAndServe(":8220", nil))
}
