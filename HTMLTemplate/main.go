package main

import (
	"html/template"
	"log"
	"net/http"
)

// html templates escapes un-safe characters avoid cross site scripting

// `<script>alert('hi')</script>`
//  in text templates, fn will be exceuted
//  in html templates , additional characters will be added in b/w

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
