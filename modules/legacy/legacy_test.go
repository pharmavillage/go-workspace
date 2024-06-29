package legacy

import (
	"testing"
)

func TestLegacy(t *testing.T) {
	result := Legacy("works")
	if result != "Legacy works" {
		t.Error("Expected Legacy to append 'works'")
	}
}
