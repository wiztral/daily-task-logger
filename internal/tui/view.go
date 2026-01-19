package tui

import (
	"daily-task-logger/internal/storage"
	"fmt"
	"strings"
)

func (m Model) View() string {
	var s strings.Builder

	// Header
	header := HeaderStyle.Render(fmt.Sprintf("%s | Total Time: %s",
		m.CurrentDate.Format("Monday, Jan 2, 2006"),
		storage.FormatMinutes(m.GetTotalTime())))
	s.WriteString(header + "\n")

	// Nav Hints (Top)
	s.WriteString(HelpStyle.Render(m.getHelpText()) + "\n")

	// Tasks
	if len(m.Tasks) == 0 {
		s.WriteString("  No tasks for this day.\n")
	} else {
		for i, task := range m.Tasks {
			cursor := " "
			style := NormalStyle
			if m.Cursor == i {
				cursor = ">"
				style = SelectedStyle
			}

			done := "[ ]"
			if task.Done {
				done = "[x]"
			}

			line := fmt.Sprintf("%s %s %s (Time: %s)", cursor, done, task.Description, storage.FormatMinutes(task.TimeMinutes))
			s.WriteString(style.Render(line) + "\n")
		}
	}

	// Input Area
	if m.Mode != ModeNav {
		s.WriteString("\n" + InputPromptStyle.Render(m.getInputPrompt()) + "\n")
		s.WriteString(m.TextInput.View() + "\n")
	}

	// Footer (Mode Status)
	s.WriteString(FooterStyle.Render(m.getFooter()))

	return s.String()
}

func (m Model) getHelpText() string {
	return "j/k: move • o: add • i: edit • d: delete • t: time • enter: toggle • y: rollover • h/l: day • g/G: top/bot • q: quit"
}

func (m Model) getInputPrompt() string {
	switch m.Mode {
	case ModeAddTask:
		return "Add Task:"
	case ModeEditTask:
		return "Edit Task:"
	case ModeEditTime:
		return "Edit Time (+15m or 1h):"
	default:
		return ""
	}
}

func (m Model) getFooter() string {
	if m.Mode == ModeNav {
		return ""
	}

	return fmt.Sprintf("\n%s", StatusInputStyle.Render("INPUT"))
}

func (m Model) getTotalTime() int {
	total := 0
	for _, t := range m.Tasks {
		total += t.TimeMinutes
	}
	return total
}
