# Runn

Naive library to handling multiple code execution and errors 

## Usage

Stop immediately if some argument is error
```go
err := runn.Execute(
    exec.Command("echo", "hello", "world"),
    SomeFunctionThatReturnError(), 
    RunnerImplementationStruct{},
)

if err != nil {
    log.Fatal(err.Error())
}
```

Run all argument while collect the errors
```go
executor := runn.Executor{
    StopWhenError: false,
}
err := executor.Execute(
    exec.Command("echo", "hello", "world"),
    SomeFunctionThatReturnError(), 
    RunnerImplementationStruct{},
)

// err is Errors type
errs, ok := err.(runn.Errors)
if ok {
    for _, er := range errs{
        fmt.Println(er.Error())
    }    
}
```

`runn.Errors` is array of error
```go
var errs runn.Errors
errs.Add(error.New("some-error"))
errs.Add(error.New("another-error"))

// aggregate error message
fmt.Println(errs.Errors())
```
