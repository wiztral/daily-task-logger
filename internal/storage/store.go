package storage

import (
	"bufio"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	BaseDir string
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
	return &Store{BaseDir: baseDir}, nil
}

func (s *Store) GetFilePath(date time.Time) string {
	filename := date.Format("2006-01-02") + ".md"
	return filepath.Join(s.BaseDir, filename)
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
