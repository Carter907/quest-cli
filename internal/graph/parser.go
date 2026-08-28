package graph

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// ParseGraph reads all markdown files in the specified directory,
// extracts their frontmatter, and returns a map of guide IDs to Guide structs.
func ParseGraph(dirPath string) (map[string]Guide, error) {
	guides := make(map[string]Guide)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {

		// ignore all other files
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		path := filepath.Join(dirPath, entry.Name())
		guide, err := parseGuide(path)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", path, err)
		}

		guides[guide.ID] = guide
	}

	return guides, nil
}

func parseGuide(path string) (Guide, error) {
	baseName := filepath.Base(path)
	id := strings.TrimSuffix(baseName, filepath.Ext(baseName))

	content, err := os.ReadFile(path)
	if err != nil {
		return Guide{}, err
	}

	lines := strings.Split(string(content), "\n")
	var frontmatterLines []string
	inFrontmatter := false

	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				break
			}
		}
		if inFrontmatter {
			frontmatterLines = append(frontmatterLines, line)
		}
	}

	if len(frontmatterLines) == 0 {
		return Guide{}, fmt.Errorf("invalid or missing YAML frontmatter")
	}

	frontmatter := strings.Join(frontmatterLines, "\n")

	var meta GuideMetadata
	if err := yaml.Unmarshal([]byte(frontmatter), &meta); err != nil {
		return Guide{}, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return Guide{
		ID:        id,
		Path:      path,
		Metadata:  meta,
		LineCount: len(lines),
	}, nil
}

// ParseConfig reads the manifest.yaml configuration file in the specified directory.
func ParseConfig(dirPath string) (Manifest, error) {
	configPath := filepath.Join(dirPath, "manifest.yaml")
	content, err := os.ReadFile(configPath)
	if err != nil {
		return Manifest{}, fmt.Errorf("failed to read manifest.yaml: %w", err)
	}

	var config Manifest
	if err := yaml.Unmarshal(content, &config); err != nil {
		return Manifest{}, fmt.Errorf("failed to parse manifest.yaml: %w", err)
	}

	return config, nil
}
