package internal

import (
	"DiffTrad/internal/config"
	"fmt"
	"os"
	"path/filepath"
)

func RetrieveFiles(c *config.Config) error {
	refLanguage, err := getLanguage(c.RefPath)
	if err != nil {
		return fmt.Errorf("Error while reading reference file: %w", err)
	}
	c.ReferenceLanguage = refLanguage

	files, err := filepath.Glob(c.TargetsPath)
	if err != nil {
		return fmt.Errorf("Error while reading target files: %w", err)
	}

	for _, file := range files {
		if file == c.RefPath {
			continue
		}

		language, err := getLanguage(file)
		if err != nil {
			return fmt.Errorf("Error while reading target file: %w", err)
		}
		c.TargetLanguages = append(c.TargetLanguages, language)
	}
	
	return nil
}

func getLanguage(filePath string) (config.Language, error) {
	dirPath, filename := filepath.Split(filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return config.Language{}, err
	}

	return config.Language{
		Content: string(content),
		Filename: filename,
		Path: dirPath,
		FullPath: filePath,
	}, nil
}