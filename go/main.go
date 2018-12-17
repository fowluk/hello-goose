package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type index_vars struct {
	Instance_id    string
	Instance_index string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	instance_id := os.Getenv("CF_INSTANCE_GUID")
	instance_index := os.Getenv("CF_INSTANCE_INDEX")
	template_file, err := ioutil.ReadFile("index.html.template")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("goose").Parse(string(template_file))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, index_vars{Instance_id: instance_id, Instance_index: instance_index})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
