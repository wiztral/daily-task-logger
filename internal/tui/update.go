package tui

import (
	"daily-task-logger/internal/storage"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Mode == ModeNav {
			return m.handleNav(msg)
		}
		return m.handleInput(msg)

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	return m, nil
}

func (m Model) handleNav(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "j", "down":
		if m.Cursor < len(m.Tasks)-1 {
			m.Cursor++
		}

	case "k", "up":
		if m.Cursor > 0 {
			m.Cursor--
		}

	case "h", "left":
		m.CurrentDate = m.CurrentDate.AddDate(0, 0, -1)
		m.Tasks, _ = m.Store.LoadTasks(m.CurrentDate)
		m.Cursor = 0

	case "l", "right":
		m.CurrentDate = m.CurrentDate.AddDate(0, 0, 1)
		m.Tasks, _ = m.Store.LoadTasks(m.CurrentDate)
		m.Cursor = 0

	case "a":
		m.Mode = ModeAddTask
		m.TextInput.Placeholder = ""
		m.TextInput.SetValue("")
		return m, textinput.Blink

	case "e":
		if len(m.Tasks) > 0 {
			m.Mode = ModeEditTask
			m.TextInput.SetValue(m.Tasks[m.Cursor].Description)
			return m, textinput.Blink
		}

	case "t":
		if len(m.Tasks) > 0 {
			m.Mode = ModeEditTime
			m.TextInput.Placeholder = ""
			m.TextInput.SetValue("")
			return m, textinput.Blink
		}

	case "d":
		if len(m.Tasks) > 0 {
			m.Tasks[m.Cursor].Done = !m.Tasks[m.Cursor].Done
			m.Store.SaveTasks(m.CurrentDate, m.Tasks)
		}

	case "x":
		if len(m.Tasks) > 0 {
			m.Tasks = append(m.Tasks[:m.Cursor], m.Tasks[m.Cursor+1:]...)
			if m.Cursor >= len(m.Tasks) && m.Cursor > 0 {
				m.Cursor--
			}
			m.Store.SaveTasks(m.CurrentDate, m.Tasks)
		}

	case "y":
		if len(m.Tasks) > 0 {
			tomorrow := m.CurrentDate.AddDate(0, 0, 1)
			tomorrowTasks, _ := m.Store.LoadTasks(tomorrow)

			taskToCopy := m.Tasks[m.Cursor]
			taskToCopy.Done = false
			taskToCopy.TimeMinutes = 0

			tomorrowTasks = append(tomorrowTasks, taskToCopy)
			m.Store.SaveTasks(tomorrow, tomorrowTasks)
		}
	}
	return m, nil
}

func (m Model) handleInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Mode = ModeNav
		return m, nil

	case "enter":
		val := m.TextInput.Value()
		if val == "" && m.Mode != ModeEditTime {
			m.Mode = ModeNav
			return m, nil
		}

		switch m.Mode {
		case ModeAddTask:
			m.Tasks = append(m.Tasks, storage.Task{
				Description: val,
				Done:        false,
				TimeMinutes: 0,
			})
		case ModeEditTask:
			m.Tasks[m.Cursor].Description = val
		case ModeEditTime:
			m.applyTimeBuffer(val)
		}

		m.Store.SaveTasks(m.CurrentDate, m.Tasks)
		m.Mode = ModeNav
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m *Model) applyTimeBuffer(input string) {
	if len(m.Tasks) == 0 {
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	prefix := ""
	if strings.HasPrefix(input, "+") {
		prefix = "+"
		input = input[1:]
	} else if strings.HasPrefix(input, "-") {
		prefix = "-"
		input = input[1:]
	}

	duration, ok := parseDuration(input)
	if !ok {
		return // Ignore invalid input to prevent reset
	}

	switch prefix {
	case "+":
		m.Tasks[m.Cursor].TimeMinutes += duration
	case "-":
		m.Tasks[m.Cursor].TimeMinutes -= duration
		if m.Tasks[m.Cursor].TimeMinutes < 0 {
			m.Tasks[m.Cursor].TimeMinutes = 0
		}
	default:
		m.Tasks[m.Cursor].TimeMinutes = duration
	}
}

func parseDuration(s string) (int, bool) {
	total := 0
	// Normalize input: "1h 3min" -> "1h 3m"
	s = strings.ReplaceAll(s, "min", "m")

	// Split into parts like ["1h", "3m"]
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return 0, false
	}

	for _, p := range parts {
		if strings.HasSuffix(p, "h") {
			val, err := strconv.Atoi(strings.TrimSuffix(p, "h"))
			if err != nil {
				return 0, false
			}
			total += val * 60
		} else if strings.HasSuffix(p, "m") {
			val, err := strconv.Atoi(strings.TrimSuffix(p, "m"))
			if err != nil {
				return 0, false
			}
			total += val
		} else {
			// Try as raw minutes
			val, err := strconv.Atoi(p)
			if err != nil {
				return 0, false
			}
			total += val
		}
	}
	return total, true
}
