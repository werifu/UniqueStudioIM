package main

import (
	"html/template"
	"log"
	"net/http"
)

func ChatroomGet(w http.ResponseWriter, r *http.Request){
	tmpl, _ := template.ParseFiles("./home.html")
	err := tmpl.Execute(w,nil)
	if err != nil{
		log.Fatal(err)
	}
}
