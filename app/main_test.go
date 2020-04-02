package main

import (
	"testing"
)

func TestMorze(t *testing.T) {
	code := translate("sos")
	if code != "... --- ..." {
		t.Errorf("Expected ... --- ... code but got %s", code)
	}
}
