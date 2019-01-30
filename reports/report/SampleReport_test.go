package report

import (
	"exercises/reports/database"
	"testing"
	"time"
)

func TestSampleReport(t *testing.T) {
	// requires a row in the database:
	// INSERT INTO member(first_name, last_name, account_balance, created_on)
	// VALUES('Some', 'Name', 0, '2013-12-11 10:09:08'::timestamp)

	report := new(SampleReport)

	rows, err := report.query(database.SampleCredentials())
	if err != nil {
		t.Error(err)
	}

	if len(rows) != 1 {
		t.Errorf("Expected %d, have %d", 1, len(rows))
	}

	if rows[0].ID != 1 {
		t.Errorf("Expected %d, have %d", 1, rows[0].ID)
	}

	if rows[0].FullName != "Some Name" {
		t.Errorf("Expected %s, have %s", "Some Name", rows[0].FullName)
	}

	if rows[0].AccountBalance != 0 {
		t.Errorf("Expected %d, have %d", 0, rows[0].AccountBalance)
	}

	if rows[0].CreatedOn.Equal(time.Date(2013, 12, 11, 10, 9, 8, 0, time.UTC)) {
		t.Errorf("Expected %s, have %s",
			time.Date(2013, 12, 11, 10, 9, 8, 0, time.UTC).Format("2006/01/02 15:04:05"),
			rows[0].CreatedOn.Format("2006/01/02 15:04:05"))
	}
}
