package drone

import (
	"testing"
)

func TestValue(t *testing.T) {
	s := ReturnDroneAsString()
	if s != "drone" {
		t.Error("Nope")
	}
}
