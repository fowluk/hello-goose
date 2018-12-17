package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type instanceInfo struct {
	InstanceID    string
	InstanceIndex string
}

var info instanceInfo

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("index.html.template")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("goose").Parse(string(f))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, info)
	if err != nil {
		panic(err)
	}
}

func main() {
	info = instanceInfo{
		InstanceID:    os.Getenv("CF_INSTANCE_GUID"),
		InstanceIndex: os.Getenv("CF_INSTANCE_INDEX"),
	}

	http.HandleFunc("/", IndexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
