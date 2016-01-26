package template

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/atamis/wt3/met/config"
	"github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template

var bufpool *bpool.BufferPool

// Load templates on program initialisation
func init() {
	bufpool = bpool.NewBufferPool(64)

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := config.Config.TemplatePath

	layouts, err := filepath.Glob(templatesDir + "layouts/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}

}

func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	// Create a buffer to temporarily write to and check if any errors were encounted.
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	fmt.Println(data)

	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}
