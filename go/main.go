package main

import (
	"net/http"
        "io/ioutil"
	"html/template"
	"os"
)

type index_vars struct {
	Hostname string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	hostname, err := os.Hostname()
	template_file, err :=  ioutil.ReadFile("index.html.template")
	if err != nil { panic(err) }
        tmpl, err := template.New("goose").Parse(string(template_file))
	if err != nil { panic(err) }
	err = tmpl.Execute(w, index_vars{Hostname: hostname})
	if err != nil { panic(err) }
}

func GooseHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("goose.jpg")
        w.Write(data)
}


func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/static/goose.jpg", GooseHandler)
	http.ListenAndServe(":8080", nil)
}
