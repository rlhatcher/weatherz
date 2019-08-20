package cmd

import "testing"

func TestMax(t *testing.T) {
	var val string
	var err error

	val, err = max("2", "6")
	if val != "6" {
		t.Errorf("Max was incorrect, got: %s, want: %s.", val, "6")
	}

	val, err = max("", "7")
	if err == nil {
		t.Errorf("Max was incorrect, got: %s, want: %s.", val, err)
	}
}

func TestMin(t *testing.T) {
	var val string
	var err error

	val, err = min("2", "6")
	if val != "2" {
		t.Errorf("Min was incorrect, got: %s, want: %s.", val, "2")
	}

	val, err = min("", "7")
	if err == nil {
		t.Errorf("Min was incorrect, got: %s, want: %s.", val, err)
	}
}
