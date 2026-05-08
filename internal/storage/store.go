package storage

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Store struct {
	BaseDir          string
	CurrentWorkspace string
}

func NewStore() (*Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	baseDir := filepath.Join(home, ".daily_task_logger")
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err = os.MkdirAll(baseDir, 0755)
		if err != nil {
			return nil, err
		}
	}

	s := &Store{
		BaseDir:          baseDir,
		CurrentWorkspace: "default",
	}

	if last, err := os.ReadFile(filepath.Join(baseDir, ".last_workspace")); err == nil {
		s.CurrentWorkspace = strings.TrimSpace(string(last))
	}

	if err := s.migrateToWorkspaces(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) migrateToWorkspaces() error {
	defaultDir := filepath.Join(s.BaseDir, "default")
	if err := os.MkdirAll(defaultDir, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(s.BaseDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".md" {
			oldPath := filepath.Join(s.BaseDir, f.Name())
			newPath := filepath.Join(defaultDir, f.Name())
			if err := os.Rename(oldPath, newPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Store) GetWorkspaceDir() string {
	workspaceDir := filepath.Join(s.BaseDir, s.CurrentWorkspace)
	_ = os.MkdirAll(workspaceDir, 0755)
	return workspaceDir
}

func (s *Store) GetFilePath(date time.Time) string {
	workspaceDir := s.GetWorkspaceDir()
	filename := date.Format("2006-01-02") + ".md"
	return filepath.Join(workspaceDir, filename)
}

func (s *Store) ListWorkspaces() []string {
	var workspaces []string
	files, err := os.ReadDir(s.BaseDir)
	if err != nil {
		return []string{"default"}
	}

	for _, f := range files {
		if f.IsDir() {
			workspaces = append(workspaces, f.Name())
		}
	}

	if len(workspaces) == 0 {
		return []string{"default"}
	}
	return workspaces
}

func (s *Store) SetWorkspace(name string) {
	if name == "" {
		return
	}
	s.CurrentWorkspace = name
	_ = os.WriteFile(filepath.Join(s.BaseDir, ".last_workspace"), []byte(name), 0644)
}

func (s *Store) LoadTasks(date time.Time) ([]Task, error) {
	path := s.GetFilePath(date)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Task{}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		task := ParseLine(line)
		if task != nil {
			tasks = append(tasks, *task)
		}
	}
	return tasks, scanner.Err()
}

func (s *Store) SaveTasks(date time.Time, tasks []Task) error {
	path := s.GetFilePath(date)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, task := range tasks {
		_, err := writer.WriteString(task.String() + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
