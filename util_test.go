package main

import (
    "testing"
)
func TestAdd(t *testing.T) {
	actual := add(2, 2);
	if (actual != 4) {
		t.Fatalf(`add(2, 2) = %q, want 4, error`, actual)
	}
}
