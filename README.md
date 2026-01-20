# Daily Task Logger

A minimalist, Go-based Terminal User Interface (TUI) for daily task logging and time tracking, featuring Vim-compliant ergonomics.

## 🚀 Features

- **Sealed Daily Logs**: Each day's tasks are stored in a dedicated Markdown file (`~/.daily_task_logger/YYYY-MM-DD.md`).
- **Minimalist TUI**: Interactive interface built with `Bubbletea` and styled with `Lipgloss`.
- **Vim-like Ergonomics**: Navigation and actions designed for Vim users.
- **Flexible Time Tracking**: Log time per task with support for absolute (`1h 20m`), additive (`+15m`), or subtractive (`-10min`) inputs.
- **Task Rollover**: Copy unfinished tasks to the next day with a single keystroke (`y`).

## ⌨️ Shortcuts

| Key | Action |
| --- | --- |
| `j` / `↓` | Move cursor down |
| `k` / `↑` | Move cursor up |
| `J` / `Shift+↓` | Move task down |
| `K` / `Shift+↑` | Move task up |
| `h` / `←` | Navigate to previous day |
| `l` / `→` | Navigate to next day |
| `o` | Add a new task (Vim "open") |
| `i` | Edit selected task description (Vim "insert") |
| `d` | Delete selected task |
| `Enter` | Toggle task completion status |
| `y` | Rollover task (yank to next day) |
| `t` | Edit logged time |
| `g` | Jump to top of list |
| `G` | Jump to bottom of list |
| `q` | Quit application |

## 🛠️ Build & Installation

### Prerequisites
- **Go** (1.18 or higher)
- **Make** (optional, recommended)

### Build
Using the provided Makefile:
```bash
make build
```
This will compile the binary to `dist/task-logger.exe`.

### Development
Run the app directly with Go:
```bash
go run main.go
```

## 📁 Storage
All data is stored in the user's home directory:
- **Windows**: `%USERPROFILE%\.daily_task_logger\`
- **Linux/macOS**: `~/.daily_task_logger/`

Files are named by date (e.g., `2024-01-19.md`) and follow standard Markdown task list syntax.
