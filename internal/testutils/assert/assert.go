package assert

import (
	"fmt"
	"testing"
)

// getErrorMessage formats the error message for assertion failures
func getErrorMessage[T any](expected, actual T) string {
	return fmt.Sprintf(
		"\nAssertion failed:\nExpected:\n\t%v\nGot:\n\t%v\n",
		expected, actual,
	)
}

// AssertEqual checks if two values are equal
func AssertEqual[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}

// AssertNotEqual checks if two values are not equal
func AssertNotEqual[T comparable](t *testing.T, actual, unexpected T) {
	t.Helper()
	if unexpected == actual {
		t.Errorf("%s", getErrorMessage(unexpected, actual))
	}
}

// AssertTrue checks if a condition is true
func AssertTrue(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Errorf("Expected condition to be true, but it was false")
	}
}

// AssertNil checks if a value is nil
func AssertNil(t *testing.T, value any) {
	t.Helper()
	if value != nil {
		t.Errorf("%s", getErrorMessage(nil, value))
	}
}

// AssertTrue checks if a condition is true
func AssertPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("%s", getErrorMessage("panic", "no panic occurred"))
		}
	}()
	f()
}

// AssertIsClose checks if a value is close to another value within a tolerance
func AssertIsClose(t *testing.T, actual, expected, tolerance float64) {
	t.Helper()
	if (expected-actual) > tolerance || (actual-expected) > tolerance {
		t.Errorf("%s", getErrorMessage(expected, actual))
	}
}
