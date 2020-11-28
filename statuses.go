package statuses

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
}

// TotalStatuses - returns the total number of statuses
func TotalStatuses(code string) int {
	return len(statuses)
}

// StatusInfo - return all info about status as Status struct
func StatusInfo(code string) *Status {
	return &Status{
		Name: statuses[code],
		Code: code,
	}
}

// AllStatuses - returns all the avialble statuses
func AllStatuses() map[string]string {
	return statuses
}
