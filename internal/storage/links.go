package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Link struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func LoadLinks(workspaceDir string) ([]Link, error) {
	path := filepath.Join(workspaceDir, "links.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Link{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var links []Link
	if err := json.Unmarshal(data, &links); err != nil {
		return nil, err
	}

	return links, nil
}

func SaveLinks(workspaceDir string, links []Link) error {
	path := filepath.Join(workspaceDir, "links.json")
	data, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
