package lint

// Rule Severity Constants
const (
	SeverityWarn  Severity = "warn"
	SeverityError Severity = "error"
)

// Severity represent the severity level of a rule
type Severity string

func (s Severity) String() string {
	switch s {
	case SeverityError:
		return "Error"
	case SeverityWarn:
		return "Warning"
	default:
		return "Severity(" + string(s) + ")"
	}
}