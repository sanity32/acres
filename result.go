package acres

import (
	"fmt"
	"io"
	"os"
)

type OutputSettings struct {
	Writer io.Writer
	Prefix string
	Muted  bool
}

var DefaultOutputSettings = OutputSettings{
	Writer: os.Stdout,
	Prefix: "Result code:",
	Muted:  false,
}

type Result struct {
	Code Code
	Err  error
	Data any
}

func (r Result) Print(v ...any) Result {
	if !DefaultOutputSettings.Muted {
		a := append([]any{"Result code:", r.Code}, v...)
		fmt.Fprintln(DefaultOutputSettings.Writer, a...)
	}
	return r
}

func (r Result) Success() bool {
	for _, code := range []Code{DONE, ALREADY, NO_JOB} {
		if r.Code == code {
			return true
		}
	}
	return false
}

func (r Result) Done() bool {
	return r.Code == DONE
}

func (r Result) Already() bool {
	return r.Code == ALREADY
}

// That includes failure
func (r Result) HasJob() bool {
	return r.Code != NO_JOB
}

func (r Result) HasResult() bool {
	return r.Code != NONE
}

func (r Result) Failure() bool {
	return r.HasResult() && !r.Success()
	// return r.Code > 0
}

func (r Result) GetCode() Code {
	return r.Code
}

func (r Result) FailureCode() Code {
	if !r.Failure() {
		return 0
	}
	return r.GetCode()
}

func (r Result) Is(codes ...Code) bool {
	for _, code := range codes {
		if r.Code == code {
			return true
		}
	}
	return false
}

func (r Result) GetError() error {
	return r.Err
}

func (r *Result) SetError(err error) *Result {
	r.Err = err
	return r
}

func (r *Result) SetData(data any) *Result {
	r.Data = data
	return r
}

func (r Result) GetData() any {
	return r.Data
}
