# acres

Action Result

## Embed usage

```
type Action struct {
    controller Controller
    Result acres.Emitter
}

func (a Action) Go() acres.Result{
    return a.Result.Done()
}
```

## Result checks

```
func DoAction(a Action){
    if r := a.Go(); r.Failure(){
        return r
    }
    return a.Result.Done()
}
```
