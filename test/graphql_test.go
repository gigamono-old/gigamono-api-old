package test

import (
	"testing"
)

// TestSample tests sample.
func TestSample(t *testing.T) {
	value := true
	if value == false {
		t.Errorf("Value is false: %v", value)
	}
}
