package fields

var TaskStatus = struct {
	InProgress string
	Completed  string
	OnHold     string
}{
	InProgress: "IN_PROGRESS",
	Completed:  "COMPLETED",
	OnHold:     "ON_HOLD",
}
