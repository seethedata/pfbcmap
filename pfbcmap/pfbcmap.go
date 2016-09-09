// pfbcmap shows the locations of Pathfinder Bootcamps
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)


type templateData  struct {
	Data string
}


func check(function string, e error) {
	if e != nil {
		log.Fatal(function, e)
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	var c http.Client
	resp,err:=c.Get(os.Getenv("dataURL"))
	check("Get Response", err)
	
	var rb []byte
	rb,err=ioutil.ReadAll(resp.Body)
	check("Read Body",err)
	
	var d templateData
	d.Data=fmt.Sprintf("%s",rb)
	var ru =r.URL.String()
	var tmpl string
	
	if  ru == "/"{
		tmpl="index.tmpl"
	} else if ru == "/alt" {
		tmpl="alt.tmpl"
	}
	
	t, err := template.ParseFiles("templates/" + tmpl)
	check("Parse template", err)
	
	t.Execute(w, d)
}



func main() {
	http.HandleFunc("/", responseHandler)
	http.HandleFunc("/alt", responseHandler)
	http.Handle("/images/", http.FileServer(http.Dir("")))
	http.Handle("/css/", http.FileServer(http.Dir("")))
	http.Handle("/fonts/", http.FileServer(http.Dir("")))
	http.Handle("/js/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
