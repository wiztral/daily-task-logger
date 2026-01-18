package tui

import (
	"daily-task-logger/internal/storage"
	"testing"
)

func TestApplyTimeBuffer(t *testing.T) {
	tests := []struct {
		initial  int
		input    string
		expected int
	}{
		{0, "30m", 30},
		{30, "+15m", 45},
		{45, "-10m", 35},
		{60, "1h 20m", 80},
		{0, "+1h", 60},
		{10, "+1min", 11},
		{5, "-10m", 0}, // No negative time
		{0, "1h 3min", 63},
		{45, "invalid", 45}, // Should not reset to 0
		{45, "+abc", 45},    // Should not reset to 0
		{45, "1h xyz", 45},  // Should not reset to 0
	}

	for _, tt := range tests {
		m := Model{
			Tasks:  []storage.Task{{TimeMinutes: tt.initial}},
			Cursor: 0,
		}
		m.applyTimeBuffer(tt.input)
		if m.Tasks[0].TimeMinutes != tt.expected {
			t.Errorf("applyTimeBuffer(%d, %q) = %d, want %d", tt.initial, tt.input, m.Tasks[0].TimeMinutes, tt.expected)
		}
	}
}
