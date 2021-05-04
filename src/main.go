package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"log"
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

	// use go rice to expose public folder
	abox, _ := rice.FindBox("public")
	if err = InitAssetsTemplates(r, abox); err != nil {
		log.Fatal(err)
	}

	// configure routes
	r.GET("/", HomeController)
	r.POST("/get-previous-prime", HighestPrimeNumber)

	// return router
	return r
}

// InitAssetsTemplates initializes the router to use the rice boxes.
// r is our main router, tbox is our template rice box, abox is our assets box
// and names are the file names of the templates to load
func InitAssetsTemplates(r *gin.Engine, abox *rice.Box) error {
	if abox != nil {
		r.StaticFS("/public", abox.HTTPBox())
	} else {
		r.Static("/public", "public")
	}
	return nil
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