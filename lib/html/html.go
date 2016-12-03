package html

import (
	"bufio"
	"io"
	"text/template"

	"github.com/Baligul/election/lib/file"
)

// Render a template by combining it with data and then output to a Writer
// Example usage:
// var b bytes.Buffer
// renderTemplate("../../templates/dynomite.service.tmpl", tmplData, &b)
func renderTemplate(tmplPath string, tmplData interface{}, w io.Writer) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, tmplData)
	if err != nil {
		return err
	}

	return nil
}

// GenerateHtmlFile function will be called in main package to automatically
// generate the html file at htmlFilePath by rendering a template
// based on the tmplPath
func GenerateHtmlFile(tmplPath string, tmplData interface{}, htmlFilePath string) error {

	// Save html file
	newFile, err := file.Create(htmlFilePath)
	defer file.Close(newFile)

	if err != nil {
		return err
	}

	wr := bufio.NewWriter(newFile)

	defer wr.Flush()

	// Render the template
	if err := renderTemplate(tmplPath, tmplData, wr); err != nil {
		return err
	}
	return nil
}
