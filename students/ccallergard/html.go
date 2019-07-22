package cyoa

import (
	"html/template"
	"os"
	"path/filepath"
)

// Generate uses given template to create HTML files based on an Adventure
func Generate(cyoa Adventure, dirPath string, templatePath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	cyoaTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	for sceneName, scene := range cyoa {
		path := filepath.Join(dirPath, sceneName+".html")
		out, err := os.Create(path)
		if err != nil {
			return err
		}
		cyoaTemplate.Execute(out, &scene)
		out.Close()
	}

	return nil

}
