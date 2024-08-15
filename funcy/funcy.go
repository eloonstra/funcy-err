package funcy

import (
	"fmt"
	"runtime"
	"strings"
)

// Err returns a new error that includes the caller's function name and the provided error.
// The caller's function name is determined using the runtime package.
func Err(err error) error {
	return fmt.Errorf("%s: %w", getCallerName(), err)
}

// getCallerName returns the name of the calling function without the package name.
func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	callerName := details.Name()
	parts := strings.Split(callerName, ".")
	return parts[len(parts)-1]
}

// ErrWithMessage returns a new error that includes the caller's function name with a message and the provided error.
// The caller's function name is determined using the runtime package.
func ErrWithMessage(err error, message string) error {
	return fmt.Errorf("%s: %s: %w", getCallerName(), message, err)
}
