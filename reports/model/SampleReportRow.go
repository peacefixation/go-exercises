package model

import "time"

// SampleReportRow a row in the Sample Report
type SampleReportRow struct {
	ID             uint64
	FullName       string
	AccountBalance int64
	CreatedOn      time.Time
}
