//easycode.go serves a webpage useful for showing a demo of Cloud Foundry
package main

import (
	"net/http"
	"html/template"
	"log"
)

func getName() string{
	return "Philadelphia, PA"
}

type Location struct {
    Name string
}

func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}


func responseHandler(w http.ResponseWriter, r *http.Request) {
	location:=Location{Name: getName() }
	t,err:=template.ParseFiles("templates/index.tmpl")
	check("Parse template",err)
	t.Execute(w,location)
}


func main() {
	http.HandleFunc("/",responseHandler)
	http.Handle("/images/",http.FileServer(http.Dir("")))
	http.Handle("/css/",http.FileServer(http.Dir("")))
	http.Handle("/fonts/",http.FileServer(http.Dir("")))
	http.Handle("/js/",http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080",nil))
}