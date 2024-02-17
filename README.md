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
