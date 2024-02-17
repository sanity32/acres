package acres

//go:generate go run golang.org/x/tools/cmd/stringer -type=Code

type Code int

const (
	DONE    Code = -1
	ALREADY Code = -2
	NO_JOB  Code = -3
	NONE    Code = 0
	FAILURE Code = 1
)
