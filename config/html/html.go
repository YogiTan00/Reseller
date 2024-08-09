package html

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"os"
)

var (
// homeTemplate    = template.Must(template.ParseFiles(filepath.Join("home", "../services/index.html")))
// serviceTemplate = template.Must(template.ParseFiles(filepath.Join("service", "../services/product/html/index.html")))
)

func InitHtml(path string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func HomeHandler() http.HandlerFunc {
	l := logger.Logger{
		EndPoint: "Html Home",
		TrxId:    uuid.New().String(),
	}
	defer l.CreateNewLog()
	tmpl, err := InitHtml(os.Getenv("html_home"))
	if err != nil {
		l.Error(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, nil)
		if err != nil {
			l.Error(err)
		}
	}
}

func ProductHandler() http.HandlerFunc {
	l := logger.Logger{
		EndPoint: "Html Product",
		TrxId:    uuid.New().String(),
	}
	defer l.CreateNewLog()
	tmpl, err := InitHtml(os.Getenv("html_product"))
	if err != nil {
		l.Error(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, nil)
		if err != nil {
			l.Error(err)
		}
	}
}
