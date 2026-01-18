package storage

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		line string
		desc string
		done bool
		time int
	}{
		{"- [ ] Simple task", "Simple task", false, 0},
		{"- [x] Done task", "Done task", true, 0},
		{"- [ ] Task with time (Time: 45m)", "Task with time", false, 45},
		{"- [x] Complex time (Time: 1h 20m)", "Complex time", true, 80},
		{"- [ ] Min suffix (Time: 10min)", "Min suffix", false, 10},
	}

	for _, tt := range tests {
		got := ParseLine(tt.line)
		if got == nil {
			t.Errorf("ParseLine(%q) = nil, want task", tt.line)
			continue
		}
		if got.Description != tt.desc || got.Done != tt.done || got.TimeMinutes != tt.time {
			t.Errorf("ParseLine(%q) = %+v, want %+v", tt.line, got, tt)
		}
	}
}

func TestFormatMinutes(t *testing.T) {
	tests := []struct {
		m    int
		want string
	}{
		{0, "0m"},
		{45, "45m"},
		{60, "1h"},
		{80, "1h 20m"},
	}

	for _, tt := range tests {
		if got := FormatMinutes(tt.m); got != tt.want {
			t.Errorf("FormatMinutes(%d) = %q, want %q", tt.m, got, tt.want)
		}
	}
}
