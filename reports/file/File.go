package file

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

// ReadQuery read a query file
func ReadQuery(filename string) (*string, error) {
	box := packr.NewBox("../sql")
	query, err := box.FindString(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read query file '%s', please recompile with 'packr build' to embed static files", filename)
	}
	return &query, nil
}

// ReadHTMLTemplate read a html template file
func ReadHTMLTemplate(filename string) (*string, error) {
	box := packr.NewBox("../html")
	query, err := box.FindString(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read html template '%s', please recompile with 'packr build' to embed static files", filename)
	}
	return &query, nil
}
