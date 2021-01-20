package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8" >
	<title>Conatenation</title>
	</head>
	<body>
	<h1>` +
		name +
		`</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error")
	}

	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}

// A defer statement defers the execution of
// a function
// until the surrounding function returns.
// The deferred call's arguments are evaluated
// immediately, but the function call is not
// executed until the surrounding function returns.
