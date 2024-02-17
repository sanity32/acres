# acres

Action Result

## Embed usage

```
type Action struct {
    controller Controller
    Result acres.Emitter
}

func (a Action) DoSomeAction() acres.Result{
    return a.Result.Done()
}
```

## Result checks

```
func (a Action)DoAnotherAction() acres.Result{
    if r := a.Go(); r.Failure(){
        return r
    }
    return a.Result.Done()
}
```

## Custom Errors

Using custom error codes:

```
//go:generate go run golang.org/x/tools/cmd/stringer -type=customError
type customError int

func (c customError) Int() int { return int(c) }

```

Declaring custom errors list:

```
const (
	minorFailure customError = iota
	generalFailure
	criticalFailure
)
// Checking our type is correct acres.ErrorCode implementation
var _ acres.ErrorCode = minorFailure
```

And later

```
func (a Action) FailyMethod() acres.Result{
    return a.Result.Err(criticalFailure)
}
``
```
