package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"strings"
)

func routerSetup() *gin.Engine {
	r := gin.Default()
	// load template files dynamically, this is helpful when building binary file
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.GET("/", HomeController)
	return r
}

// load template as proveded in go-assets builder documentation to load templates, and provide them in binary build
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func main() {
	r := routerSetup()
	r.Run(":8080")
}