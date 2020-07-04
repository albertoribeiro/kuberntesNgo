package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	message := greeting("Code.education Rocks !!!") //"<b>Code.education Rocks!</b>"
	data := map[string]string{
		"Message": message,
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}
