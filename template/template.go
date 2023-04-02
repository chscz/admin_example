package template

import (
	"errors"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

const (
	viewPath     = "web_content/view/"
	templatePath = viewPath + "template/"
)

var templateList = []CustomTemplate{
	{"home", true},
	{"about", true},
	{"user_list", true},
}

type CustomTemplate struct {
	Name    string
	UseBase bool
}

func SetTemplate() *TemplateRegistry {
	templates := make(map[string]*template.Template)
	for _, temp := range templateList {
		if temp.UseBase {
			templates[temp.Name] = template.Must(template.ParseFiles(
				templatePath+temp.Name+".html",
				viewPath+"base.html",
			))
		} else {
			templates[temp.Name] = template.Must(template.ParseFiles(
				templatePath + temp.Name + ".html",
			))
		}
	}
	return &TemplateRegistry{
		templates: templates,
	}
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
