package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestContainsOnly(t *testing.T) {
	chars := "ABCDEF"

	if !str.ContainsOnly("ABC", chars) {
		t.Error("ContainsOnly MUST return true for ABC contains only chars: " + chars)
	}

	if str.ContainsOnly("XYZ", chars) {
		t.Error("ContainsOnly MUST return false for XYZ contains only chars: " + chars)
	}

	if str.ContainsOnly("XAZ", chars) {
		t.Error("ContainsOnly MUST return false for XAZ contains only chars: " + chars)
	}
}
