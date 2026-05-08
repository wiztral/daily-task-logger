# Daily Task Logger

A minimalist, Go-based Terminal User Interface (TUI) for daily task logging and time tracking, featuring Vim-compliant ergonomics.

## 🚀 Features

- **Workspaces**: Organize tasks into project-level folders (e.g., `Work`, `Personal`).
- **Sealed Daily Logs**: Each day's tasks are stored in a dedicated Markdown file within the workspace folder.
- **Minimalist TUI**: Interactive interface built with `Bubbletea` and styled with `Lipgloss`.
- **Vim-like Ergonomics**: Navigation and actions designed for Vim users.
- **Flexible Time Tracking**: Log time per task with support for absolute (`1h 20m`), additive (`+15m`), or subtractive (`-10min`) inputs.
- **Task Rollover**: Copy unfinished tasks to the next day with a single keystroke (`y`).
- **External Task Links**: Manage, fuzzy-search, and attach external links (like Notion docs or Jira tickets) directly to your tasks, complete with browser integration.

## ⌨️ Shortcuts

| Key | Action |
| --- | --- |
| `w` | Switch or create Workspace |
| `[` / `]` | Cycle through Workspaces |
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
| `a` | Add a new external link to workspace store |
| `s` | Fuzzy search links and attach to selected task |
| `S` | Fuzzy search links and create new task from link |
| `v` | Open attached link of selected task in browser |
| `g` | Jump to top of list |
| `G` | Jump to bottom of list |
| `q` | Quit application |

*(When fuzzy-searching links, you can use `Ctrl+P` and `Ctrl+N` to navigate up and down without exiting the filter input).*

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
All data is stored in the user's home directory under workspace subdirectories:
- **Windows**: `%USERPROFILE%\.daily_task_logger\<workspace>\`
- **Linux/macOS**: `~/.daily_task_logger/<workspace>/`

Daily logs are named by date (e.g., `2024-01-19.md`) and follow standard Markdown task list syntax. Links attached to tasks are saved natively in the Markdown (e.g. `- [ ] [Task Title](URL)`). 

External links for the workspace are saved in a standard `links.json` file inside the workspace directory. Existing unorganized tasks are automatically migrated to a `default` workspace on first run.
