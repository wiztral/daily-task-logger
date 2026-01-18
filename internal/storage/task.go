package storage

import (
	"fmt"
	"regexp"
	"strings"
)

type Task struct {
	Done        bool
	Description string
	TimeMinutes int
}

var taskRegex = regexp.MustCompile(`^- \[( |x)\] (.*?) \(Time: (.*?)\)`)
var simpleRegex = regexp.MustCompile(`^- \[( |x)\] (.*)`)

func ParseLine(line string) *Task {
	matches := taskRegex.FindStringSubmatch(line)
	if len(matches) == 4 {
		return &Task{
			Done:        matches[1] == "x",
			Description: matches[2],
			TimeMinutes: parseTime(matches[3]),
		}
	}

	matches = simpleRegex.FindStringSubmatch(line)
	if len(matches) == 3 {
		return &Task{
			Done:        matches[1] == "x",
			Description: matches[2],
			TimeMinutes: 0,
		}
	}

	return nil
}

func parseTime(s string) int {
	total := 0
	// Normalize "min" to "m"
	s = strings.ReplaceAll(s, "min", "m")

	parts := strings.Fields(s)
	for _, p := range parts {
		if strings.HasSuffix(p, "h") {
			var h int
			fmt.Sscanf(p, "%dh", &h)
			total += h * 60
		} else if strings.HasSuffix(p, "m") {
			var m int
			fmt.Sscanf(p, "%dm", &m)
			total += m
		}
	}
	return total
}

func (t Task) String() string {
	doneStr := " "
	if t.Done {
		doneStr = "x"
	}
	return fmt.Sprintf("- [%s] %s (Time: %s)", doneStr, t.Description, FormatMinutes(t.TimeMinutes))
}

func FormatMinutes(m int) string {
	if m == 0 {
		return "0m"
	}
	h := m / 60
	remM := m % 60
	if h > 0 && remM > 0 {
		return fmt.Sprintf("%dh %dm", h, remM)
	} else if h > 0 {
		return fmt.Sprintf("%dh", h)
	}
	return fmt.Sprintf("%dm", remM)
}
