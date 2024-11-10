package store

import (
	"fmt"
	"netsuite-companion/util"
	"os"
	"path/filepath"
)

// CreateGlobal creates a new global store file
func (s *BaseStore) CreateGlobal(force bool) error {
	// Get the path for the global store file
	path, err := s.getGlobalPath()
	if err != nil {
		return err
	}

	// Check if the file exists and if force is true, or if the file does not exist
	if !util.Exists(path) || force {
		// Collect input for the global store
		store, err := s.collectGlobalInput()
		if err != nil {
			return err
		}

		// Save the global store to the file
		if err := s.saveToFile(path, store); err != nil {
			return err
		}
	}

	// If no error occurred, return nil
	return nil
}

// RetrieveGlobal retrieves the global store from the file
func (s *BaseStore) RetrieveGlobal() (*GlobalStore, error) {
	// Get the path for the global store file
	path, err := s.getGlobalPath()
	if err != nil {
		return nil, err
	}

	// Read the global store file
	store, err := s.readGlobalFile(path)
	if err != nil {
		return nil, err
	}

	// If no error occurred, return the global store
	return store, nil
}

// UpdateGlobal updates the global store in the file
func (s *BaseStore) UpdateGlobal(store *GlobalStore) error {
	// Get the path for the global store file
	path, err := s.getProjectPath()
	if err != nil {
		return err
	}

	// Update the global store file
	err = s.updateFile(path, store)
	if err != nil {
		return err
	}

	// If no error occurred, return nil
	return nil
}

// collectGlobalInput prompts the user for store configuration values
func (s *BaseStore) collectGlobalInput() (*GlobalStore, error) {
	// Create a new global store
	store := GlobalStore{}

	// Get the author name from the user
	store.AuthorName = util.GetInput("Enter author name:")
	if store.AuthorName == "" {
		return nil, fmt.Errorf("author name must be non-empty")
	}

	// Get the author email from the user
	store.AuthorEmail = util.GetInput("Enter author email:")
	if store.AuthorEmail == "" {
		return nil, fmt.Errorf("author email must be non-empty")
	}

	// Get the vendor name from the user
	store.VendorName = util.GetInput("Enter vendor name:")
	if store.VendorName == "" {
		return nil, fmt.Errorf("vendor name must be non-empty")
	}

	// Get the vendor prefix from the user
	store.VendorPrefix = util.GetInput("Enter three letter vendor prefix:")
	if store.VendorPrefix == "" {
		return nil, fmt.Errorf("vendor prefix must be non-empty")
	}

	// Validate the vendor prefix
	err := s.validateVendorPrefix(store.VendorPrefix)
	if err != nil {
		return nil, err
	}

	// Return the global store
	return &store, nil
}

// getGlobalPath returns the full path to the global store file
func (s *BaseStore) getGlobalPath() (string, error) {
	// Get the user's home directory
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Join the directory name with the store directory name
	return filepath.Join(dirname, storeDirName), nil
}

// validateVendorPrefix ensures the vendor prefix is exactly 3 characters
func (s *BaseStore) validateVendorPrefix(prefix string) error {
	// Check if the length of the vendor prefix is exactly 3
	if len(prefix) != 3 {
		return fmt.Errorf("vendor prefix must be exactly 3 characters")
	}

	// If the length is 3, return nil
	return nil
}
