## funcy-err

`funcy` is a Go package that provides a simple and intuitive way to wrap errors with the caller's function name and custom messages.

### Installation

To install the `funcy` package, you can use `go get`:

```sh
go get github.com/eloonstra/funcy-err
```

### Functions

#### Err

Wraps an existing error with the caller's function name.

```go
func Err(err error) error
```

#### ErrWithMessage

Wraps an existing error with the caller's function name and a custom message.

```go
func ErrWithMessage(err error, message string) error
```

### Internal Functions

#### getCallerName

Retrieves the name of the calling function. This function is used internally to get the caller's function name without the package name.

```go
func getCallerName() string
```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.