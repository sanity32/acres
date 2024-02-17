package acres

type FuncWithResult func() Result

type FuncWithError func() error

func RunNonFailureResultChain(ff ...FuncWithResult) (lastResult Result) {
	for _, f := range ff {
		if lastResult = f(); lastResult.Failure() {
			break
		}
	}
	return
}

func RunNonFailureErrorChain(successResult Result, ff ...FuncWithError) (lastResult Result) {
	for _, f := range ff {
		if f == nil {
			continue
		}
		if err := f(); err != nil {
			return NewEmitter().Failure(err)
		}
	}
	return successResult
}

func RunResultyChain(ff ...FuncWithResult) (r Result) {
	for _, f := range ff {
		if f != nil {
			if r = f(); r.HasResult() {
				break
			}
		}
	}
	return r
}
