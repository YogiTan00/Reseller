package html

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"github.com/google/uuid"
	"html/template"
	"log"
	"net/http"
)

const (
	HtmlPath = "HTML_PATH"
)

var (
// homeTemplate    = template.Must(template.ParseFiles(filepath.Join("home", "../services/index.html")))
// serviceTemplate = template.Must(template.ParseFiles(filepath.Join("service", "../services/product/html/index.html")))
)

type Config struct {
	HTMLPath string
}

func NewConfig() *Config {
	return &Config{
		HTMLPath: utils.NewEnv(HtmlPath),
	}
}

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
	cfg := NewConfig()
	tmpl, err := InitHtml(cfg.HTMLPath + "/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ProductHandler() http.HandlerFunc {
	l := logger.Logger{
		EndPoint: "Html Product",
		TrxId:    uuid.New().String(),
	}
	defer l.CreateNewLog()
	cfg := NewConfig()
	tmpl, err := InitHtml(cfg.HTMLPath + "/product/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TransactionHandler() http.HandlerFunc {
	l := logger.Logger{
		EndPoint: "Html Transaction",
		TrxId:    uuid.New().String(),
	}
	defer l.CreateNewLog()
	cfg := NewConfig()
	tmpl, err := InitHtml(cfg.HTMLPath + "/transaction/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
