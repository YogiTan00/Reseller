package html

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"github.com/google/uuid"
	"html/template"
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
	tmpl, err := InitHtml(cfg.HTMLPath + "/services/html/index.html")
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
	cfg := NewConfig()
	tmpl, err := InitHtml(cfg.HTMLPath + "/services/html/product/index.html")
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
