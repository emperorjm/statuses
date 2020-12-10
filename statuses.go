package statuses

import (
	"errors"
)

// Status - all info about status
type Status struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var statuses = map[string]string{
	"active":            "Active",
	"approved":          "Approved",
	"awaiting_approval": "Awaiting Approval",
	"blocked":           "blocked",
	"cancelled":         "Cancelled",
	"closed":            "Closed",
	"completed":         "Completed",
	"disabled":          "Disabled",
	"disapproved":       "Disapproved",
	"hold":              "Hold",
	"in_progress":       "In Progress",
	"inactive":          "Inactive",
	"moderation":        "Moderation",
	"open":              "Open",
	"pending":           "Pending",
	"processed":         "Processed",
	"processing":        "Processing",
	"refund":            "Refund",
	"rejected":          "Rejected",
	"success":           "Success",
	"verified":          "Verified",
}

// TotalStatuses - returns the total number of statuses
func TotalStatuses() int {
	return len(statuses)
}

// StatusInfo - return all info about status as Status struct
func StatusInfo(code string) (*Status, error) {
	// Check if the map key exists and if not returns an error with a message.
	name, ok := statuses[code]
	if !ok {
		return nil, errors.New("Invalid code submitted")
	}

	return &Status{
		Name: name,
		Code: code,
	}, nil
}

// AllStatuses - returns all the avialble statuses
func AllStatuses() map[string]string {
	return statuses
}
