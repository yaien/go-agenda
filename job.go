package agenda

// Job structure
type Job struct {
	UUID         string
	Name         string
	FailedAt     string
	FailedReason string
	CompletedAt  string
	NextRun      string
	Data         interface{}
}
