package template

import (
	"bytes"
	"exercises/reports/file"
	"exercises/reports/math"
	"fmt"
	"html/template"
	"time"
)

// ProcessTemplate process a template file with the given data
func ProcessTemplate(templateFile, reportTitle string, data interface{}) (string, error) {
	templateString, err := file.ReadHTMLTemplate(templateFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	t := template.Must(template.New(reportTitle).Funcs(templateFuncs).Parse(*templateString))
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// templateFuncs functions to pass to the HTML template processor
var templateFuncs = template.FuncMap{
	"formatDate": func(date time.Time, dateFormat string) string {
		return date.Format(dateFormat)
	},
	"formatCurrency": func(cents int64) string {
		sign := ""
		if cents < 0 {
			sign = "-"
		}

		abs := math.Abs(cents)
		dollarsPart := abs / 100
		centsPart := abs % 100

		return fmt.Sprintf("%s$%d.%02d", sign, dollarsPart, centsPart)
	},
}
