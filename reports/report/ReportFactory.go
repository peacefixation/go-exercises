package report

import (
	"fmt"
	"strings"
)

// Factory returns an instance of a report
type Factory func() Report

var factories = make(map[string]Factory)

// Register a report factory
func Register(reportName string, factory Factory) {
	factories[reportName] = factory
}

// Create the report with the specified name
func Create(reportName string) (Report, error) {
	reportFactory, ok := factories[reportName]
	if !ok {
		// report factory is not registered, show available reports
		availableReports := make([]string, 0)
		for k := range factories {
			availableReports = append(availableReports, k)
		}
		return nil, fmt.Errorf("Invalid report name '%s'. Available reports: '%s'",
			reportName,
			strings.Join(availableReports, ", "))
	}
	return reportFactory(), nil
}
