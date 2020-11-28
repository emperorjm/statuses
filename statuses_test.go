package statuses

import (
	"testing"
)

// TestActiveCode tries to get the status info with
// code 'active'
func TestActiveCode(t *testing.T) {
	code := "active"
	status, err := StatusInfo(code)
	if err != nil {
		t.Fatalf(`%v is not a valid code`, code)
	} else {
		t.Logf("Status: %v\n", *status)
	}
}

// TestInvalidCode tries to get the status info with
// an invalid code 'invalid'
func TestInvalidCode(t *testing.T) {
	code := "invalid"
	_, err := StatusInfo(code)
	if err == nil {
		t.Fatalf(`%v code is actually valid`, code)
	}
}

// TestTotalStatuses get the total number of statuses
// that are being stored
func TestTotalStatuses(t *testing.T) {
	count := TotalStatuses()
	t.Logf("Total Statuses: %v\n", count)
}

// TestGetAllStatuses get all the statuses
// that are being stored
func TestAllStatuses(t *testing.T) {
	statuses := AllStatuses()
	count := TotalStatuses()

	if len(statuses) != count {
		t.Fatalf("Invalid number of statuses returned")
	} else {
		t.Logf("Statuses: %v\n", statuses)
	}
}
