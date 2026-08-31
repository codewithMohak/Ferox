package engine

import (
	"testing"
)

func TestProcessJobSafely_RecoversPanic(t *testing.T) {
	results := make(chan Result, 1)

	job := Job{
		URL: "panic://test",
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				results <- Result{
					Job: job,
					Err: panicToError(r),
				}
			}
		}()

		panic("test panic")
	}()

	result := <-results

	if result.Err == nil {
		t.Fatal("expected panic to become an error")
	}
}

func panicToError(r interface{}) error {
	switch value := r.(type) {
	case error:
		return value
	default:
		return &panicError{value: value}
	}
}

type panicError struct {
	value interface{}
}

func (e *panicError) Error() string {
	return "panic: " + formatPanicValue(e.value)
}

func formatPanicValue(value interface{}) string {
	return stringValue(value)
}

func stringValue(value interface{}) string {
	if s, ok := value.(string); ok {
		return s
	}

	return "unknown panic value"
}
