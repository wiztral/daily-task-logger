package tui

import (
	"daily-task-logger/internal/storage"
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Mode int

const (
	ModeNav Mode = iota
	ModeAddTask
	ModeEditTask
	ModeEditTime
	ModeWorkspace
	ModeSearchLinks
	ModeAddLinkTitle
	ModeAddLinkURL
)

const (
	IntentAttach = iota
	IntentCreate
)

type linkItem struct {
	link storage.Link
}

func (i linkItem) Title() string       { return i.link.Title }
func (i linkItem) Description() string { return i.link.URL }
func (i linkItem) FilterValue() string { return i.link.Title + " " + i.link.URL }

type linkDelegate struct{}

func (d linkDelegate) Height() int                               { return 1 }
func (d linkDelegate) Spacing() int                              { return 0 }
func (d linkDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d linkDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(linkItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s (%s)", i.link.Title, i.link.URL)
	fn := NormalStyle.Render
	cursor := " "

	if index == m.Index() {
		fn = SelectedStyle.Render
		cursor = ">"
	}

	fmt.Fprintf(w, "%s %s", cursor, fn(str))
}

type Model struct {
	Store        *storage.Store
	CurrentDate  time.Time
	Tasks        []storage.Task
	Cursor       int
	Mode         Mode
	TextInput    textinput.Model
	LinkList     list.Model
	LinkInput    textinput.Model
	LinkIntent   int
	NewLinkTitle string
	Workspace    string
	Error        error
	Width        int
	Height       int
}

func NewModel(s *storage.Store) Model {
	ti := textinput.New()
	ti.Focus()
	ti.PlaceholderStyle = PlaceholderStyle

	li := textinput.New()
	li.Focus()
	li.PlaceholderStyle = PlaceholderStyle

	ll := list.New([]list.Item{}, linkDelegate{}, 0, 0)
	ll.Title = "Search Links"
	ll.SetShowStatusBar(false)

	ll.KeyMap.CursorUp = key.NewBinding(
		key.WithKeys("up", "k", "ctrl+p"),
		key.WithHelp("↑/k/ctrl+p", "up"),
	)
	ll.KeyMap.CursorDown = key.NewBinding(
		key.WithKeys("down", "j", "ctrl+n"),
		key.WithHelp("↓/j/ctrl+n", "down"),
	)

	now := time.Now()
	tasks, _ := s.LoadTasks(now)

	return Model{
		Store:       s,
		CurrentDate: now,
		Tasks:       tasks,
		Mode:        ModeNav,
		TextInput:   ti,
		LinkInput:   li,
		LinkList:    ll,
		Workspace:   s.CurrentWorkspace,
	}
}

func (m Model) GetTotalTime() int {
	total := 0
	for _, t := range m.Tasks {
		total += t.TimeMinutes
	}
	return total
}
