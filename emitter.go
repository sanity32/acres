package acres

import "errors"

type Emitter struct{}

func NewEmitter() Emitter {
	return Emitter{}
}

func (g Emitter) Code(code Code, err ...error) Result {
	var e error
	if a := err; len(a) > 0 {
		e = err[0]
	}
	return Result{
		Code: code,
		Err:  e,
	}
}

func (g Emitter) None(err ...error) Result {
	return g.Code(NONE, err...)
}

// Returns DONE = -1
func (g Emitter) Done(data ...any) Result {
	r := g.Code(DONE)
	if a := data; len(a) > 0 {
		r.SetData(a[0])
	}
	return r
}

// Returns ALREADY = -2
func (g Emitter) Already(err ...error) Result {
	return g.Code(ALREADY, err...)
}

// Returns NO_JOB = -3
func (g Emitter) NoJob(err ...error) Result {
	return g.Code(NO_JOB, err...)
}

// Returns general general FAILURE = -1
func (g Emitter) Failure(err ...error) Result {
	return g.Code(FAILURE, err...)
}

func (g Emitter) Err(code ErrorCode) Result {
	return g.Code(Code(code.Int()), errors.New(code.String()))
}

func (g Emitter) FromErr(err error, result ...Result) Result {
	return g.FromBool(err == nil, result...)
}

func (g Emitter) FromBool(success bool, result ...Result) Result {
	if success {
		if a := result; len(a) > 0 {
			return a[0]
		}
		return g.Done()
	}
	return g.Failure()
}
