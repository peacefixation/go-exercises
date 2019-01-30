package report

import (
	"exercises/reports/database"
	"exercises/reports/file"
	"exercises/reports/model"
)

// SampleReport report on members
type SampleReport struct {
	title      string
	content    string
	recipients []string
}

// NewSampleReport construct an instance of the Payments report
func NewSampleReport() Report {
	return &SampleReport{
		title: "SampleReport",
	}
}

// register the report
func init() {
	Register("SampleReport", NewSampleReport)
}

// Title the title of the report
func (report SampleReport) Title() string {
	return report.title
}

// Run the report
func (report SampleReport) Run(dbCredentials database.Credentials, queryArgs ...interface{}) (*string, error) {
	rows, err := report.query(dbCredentials, queryArgs...)
	if err != nil {
		return nil, err
	}

	return processTemplate(report.title, rows)
}

// Query execute the query and return the rows
func (report SampleReport) query(dbCredentials database.Credentials, args ...interface{}) ([]*model.SampleReportRow, error) {
	queryString, err := file.ReadQuery("SampleReport.sql")
	if err != nil {
		return nil, err
	}

	db, err := database.New(dbCredentials)
	if err != nil {
		return nil, err
	}

	sqlRows, err := db.Query(*queryString, args...)
	if err != nil {
		return nil, err
	}
	defer sqlRows.Close()

	reportRows := make([]*model.SampleReportRow, 0)

	for sqlRows.Next() {
		row := new(model.SampleReportRow)
		err = sqlRows.Scan(&row.ID, &row.FullName, &row.AccountBalance, &row.CreatedOn)
		if err != nil {
			return nil, err
		}
		reportRows = append(reportRows, row)
	}

	return reportRows, nil
}
