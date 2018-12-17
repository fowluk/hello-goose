package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type vars struct {
	InstanceID    string
	InstanceIndex string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	id := os.Getenv("CF_INSTANCE_GUID")
	index := os.Getenv("CF_INSTANCE_INDEX")
	f, err := ioutil.ReadFile("index.html.template")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("goose").Parse(string(f))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, vars{InstanceID: id, InstanceIndex: index})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
