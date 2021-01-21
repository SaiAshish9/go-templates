package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type a1 struct {
	A string
	B int
}

type b1 struct {
	C []a1
}

func init() {
	tpl = template.Must(template.ParseGlob("*"))
	// templates/*.txt
	// templates/*
}

func main() {

	// every type implements the empty interface

	err := tpl.Execute(os.Stdout, nil)
	var data interface{}
	data = []string{"a", "b", "c"}
	data = map[string]string{
		"a": "1"}
	data = a1{A: "a", B: 1}
	data = b1{
		C []{data,data}
	}
	_ = data


	var fm = template.FuncMap(
		"us": strings.ToUpper,
		"ft": firstThree
	)

	// err = tpl.ExecuteTemplate(os.Stdout, "a.txt", nil)
	// err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	// err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", `test`)

	err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}

}
