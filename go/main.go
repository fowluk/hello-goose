package main

import (
	"html/template"
	"net/http"
	"os"
)

type instanceInfo struct {
	InstanceID    string
	InstanceIndex string
}

var (
	info instanceInfo
	tmpl *template.Template
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, info); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	info = instanceInfo{
		InstanceID:    os.Getenv("CF_INSTANCE_GUID"),
		InstanceIndex: os.Getenv("CF_INSTANCE_INDEX"),
	}

	tmpl = template.Must(template.New("goose").Parse(gooseTemplate))

	http.HandleFunc("/", IndexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}
