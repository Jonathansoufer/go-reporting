package reportstore

import "time"

const (
	ReportStatusUnknown ReportStatus = iota
	ReportStatusPending
	ReportStatusReviewing
	ReportStatusResolved
)

func (s ReportStatus) String() string {
	switch s {
	case ReportStatusPending:
		return "pending"
	case ReportStatusReviewing:
		return "reviewing"
	case ReportStatusResolved:
		return "resolved"
	default:
		return ""
	}
}

func ParseReportStatus(s string) ReportStatus {
	switch s {
	case "Pending":
		return ReportStatusPending
	case "Reviewing":
		return ReportStatusReviewing
	case "Resolved":
		return ReportStatusResolved
	default:
		return ReportStatusUnknown
	}
}

type Report struct {
	ReportID string `json:"report_id"`
	UserID string `json:"user_id"`
	ResolverID string `json:"resolver_id"`
	ReviewerIDs []string `json:"reviewer_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status	ReportStatus
	Title string `json:"title"`
}

type ReportStatus int

type Message struct {
	ReportID string `json:"report_id"`
	MessageID string `json:"message_id"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}