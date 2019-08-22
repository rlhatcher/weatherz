package cmd

import "testing"

func TestStrMax(t *testing.T) {
	var val string
	var err error

	val, err = strmax("2", "6")
	if val != "6" {
		t.Errorf("Max was incorrect, got: %s, want: %s.", val, "6")
	}

	val, err = strmax("", "7")
	if err == nil {
		t.Errorf("Max was incorrect, got: %s, want: %s.", val, err)
	}
}

func TestStrMin(t *testing.T) {
	var val string
	var err error

	val, err = strmin("2", "6")
	if val != "2" {
		t.Errorf("Min was incorrect, got: %s, want: %s.", val, "2")
	}

	val, err = strmin("", "7")
	if err == nil {
		t.Errorf("Min was incorrect, got: %s, want: %s.", val, err)
	}
}
