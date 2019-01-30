package report

import (
	"exercises/reports/database"
	"exercises/reports/template"
	t "html/template"
	"time"
)

// Report the interface that each report should implement
type Report interface {
	// Run the report
	Run(dbCredentials database.Credentials, queryArgs ...interface{}) (*string, error)
	// Title the title of the report
	Title() string
}

// ProcessTemplate the report template
func processTemplate(reportTitle string, rows interface{}) (*string, error) {

	body, err := template.ProcessTemplate("SampleReport.html", reportTitle, rows)
	if err != nil {
		return nil, err
	}

	header, err := template.ProcessTemplate("Header.html", reportTitle, reportTitle)
	if err != nil {
		return nil, err
	}

	footer, err := template.ProcessTemplate("Footer.html", reportTitle, time.Now())
	if err != nil {
		return nil, err
	}

	type parts struct {
		Body   t.HTML
		Header t.HTML
		Footer t.HTML
	}

	report, err := template.ProcessTemplate(
		"Wrapper.html",
		reportTitle,
		parts{
			Body:   t.HTML(body),
			Header: t.HTML(header),
			Footer: t.HTML(footer),
		},
	)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
