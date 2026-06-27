package engine

import "testing"

func TestColorize_Enabled(t *testing.T) {
	orig := colorEnabled
	colorEnabled = true
	t.Cleanup(func() { colorEnabled = orig })

	got := colorize(Red, "hello")
	want := Red + "hello" + Reset
	if got != want {
		t.Errorf("colorize enabled: got %q, want %q", got, want)
	}
}

func TestColorize_Disabled(t *testing.T) {
	orig := colorEnabled
	colorEnabled = false
	t.Cleanup(func() { colorEnabled = orig })

	got := colorize(Red, "hello")
	if got != "hello" {
		t.Errorf("colorize disabled: got %q, want %q", got, "hello")
	}
}
