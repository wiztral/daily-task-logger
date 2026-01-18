package tui

import (
	"daily-task-logger/internal/storage"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type Mode int

const (
	ModeNav Mode = iota
	ModeAddTask
	ModeEditTask
	ModeEditTime
)

type Model struct {
	Store       *storage.Store
	CurrentDate time.Time
	Tasks       []storage.Task
	Cursor      int
	Mode        Mode
	TextInput   textinput.Model
	Error       error
	Width       int
	Height      int
}

func NewModel(s *storage.Store) Model {
	ti := textinput.New()
	ti.Focus()
	ti.PlaceholderStyle = PlaceholderStyle

	now := time.Now()
	tasks, _ := s.LoadTasks(now)

	return Model{
		Store:       s,
		CurrentDate: now,
		Tasks:       tasks,
		Mode:        ModeNav,
		TextInput:   ti,
	}
}

func (m Model) GetTotalTime() int {
	total := 0
	for _, t := range m.Tasks {
		total += t.TimeMinutes
	}
	return total
}
