package store

import (
	"netsuite-companion/util"
	"os"
	"path/filepath"
)

// CreateProject creates a new project
func (s *BaseStore) CreateProject() error {
	// Get the path for the project file
	path, err := s.getProjectPath()
	if err != nil {
		return err
	}

	// Collect input for the project
	store, err := s.collectProjectInput()
	if err != nil {
		return err
	}

	// Save the project to the file
	if err := s.saveToFile(path, store); err != nil {
		return err
	}

	// If no error occurred, return nil
	return nil
}

// RetrieveProject retrieves a project
func (s *BaseStore) RetrieveProject() (*ProjectStore, error) {
	// Get the path for the project file
	path, err := s.getProjectPath()
	if err != nil {
		return nil, err
	}

	// Read the project file
	store, err := s.readProjectFile(path)
	if err != nil {
		return nil, err
	}

	// If no error occurred, return the project
	return store, nil
}

// UpdateProject updates a project
func (s *BaseStore) UpdateProject(store *ProjectStore) error {
	// Get the path for the project file
	path, err := s.getProjectPath()
	if err != nil {
		return err
	}

	// Update the project file
	err = s.updateFile(path, store)
	if err != nil {
		return err
	}

	// If no error occurred, return nil
	return nil
}

// getProjectPath gets the path for the project file
func (s *BaseStore) getProjectPath() (string, error) {
	// Get the current working directory
	dirname, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Join the directory name with the store directory name
	return filepath.Join(dirname, storeDirName), nil
}

// collectProjectInput collects input for the project
func (s *BaseStore) collectProjectInput() (*ProjectStore, error) {
	// Create a new project store
	store := &ProjectStore{}

	// Get the project name from the user
	store.Current = util.GetInput("Enter project name: ")

	// Return the project store
	return store, nil
}
