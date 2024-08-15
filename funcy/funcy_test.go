package funcy

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestErr(t *testing.T) {
	err := errors.New("test error")
	want := fmt.Errorf("TestErr: %w", err)

	if got := Err(err); !reflect.DeepEqual(got, want) {
		t.Errorf("Err() = %v, want %v", got, want)
	}
}

func TestErrWithMessage(t *testing.T) {
	err := errors.New("test error")
	message := "test message"
	want := fmt.Errorf("TestErrWithMessage: %s: %w", message, err)

	if got := ErrWithMessage(err, message); !reflect.DeepEqual(got, want) {
		t.Errorf("ErrWithMessage() = %v, want %v", got, want)
	}
}
