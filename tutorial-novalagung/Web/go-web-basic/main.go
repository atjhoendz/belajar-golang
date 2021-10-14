package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var templ, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]string{
		"title": "Learning Golang Web",
		"name":  "Gotham",
	}

	err = templ.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var address = "localhost:3000"

	fmt.Printf("Server started at %s\n", address)

	server := new(http.Server)

	server.Addr = address

	err := server.ListenAndServe()

	// err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
