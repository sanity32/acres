package acres

type customError int

func (c customError) Int() int { return int(c) }

//go:generate go run golang.org/x/tools/cmd/stringer -type=customError
const (
	minorFailure customError = iota
	generalFailure
	criticalFailure
)

var _ ErrorCode = minorFailure
