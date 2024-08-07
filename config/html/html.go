package html

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	homeTemplate    = template.Must(template.ParseFiles(filepath.Join("home", "../services/index.html")))
	serviceTemplate = template.Must(template.ParseFiles(filepath.Join("service", "../services/product/html/index.html")))
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := homeTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	err := serviceTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
