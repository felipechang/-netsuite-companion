package store

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"netsuite-companion/util"
	"os"
)

// Constants
const (
	storeDirName     = ".nsc" // Directory name for storing configuration
	storePermissions = 0600   // File permissions (read/write for owner only)
)

// Store interface
type Store interface {
	CreateGlobal(force bool) error
	CreateProject() error
	RetrieveGlobal() (*GlobalStore, error)
	RetrieveProject() (*ProjectStore, error)
	UpdateGlobal(store *GlobalStore) error
	UpdateProject(store *ProjectStore) error
}

type ProjectStore struct {
	Current string `yaml:"current"`
}

type GlobalStore struct {
	AuthorName   string `yaml:"author_name"`
	AuthorEmail  string `yaml:"author_email"`
	VendorName   string `yaml:"vendor_name"`
	VendorPrefix string `yaml:"vendor_prefix"`
	OpenAIApiKey string `yaml:"openai_api_key"`
}

type BaseStore struct {
}

// NewBaseStore creates a new BaseStore
func NewBaseStore() *BaseStore {
	return &BaseStore{}
}

// saveToFile saves content to a file
func (s *BaseStore) saveToFile(path string, content interface{}) error {
	// Open the file for writing
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, storePermissions)
	if err != nil {
		return err
	}
	defer file.Close()
	// Encode the content to YAML and write it to the file
	if err := yaml.NewEncoder(file).Encode(content); err != nil {
		return err
	}
	return nil
}

// readProjectFile reads a project file
func (s *BaseStore) readProjectFile(path string) (*ProjectStore, error) {
	// Check if the file exists
	if !util.Exists(path) {
		return nil, fmt.Errorf("store not found at %s", path)
	}
	// Open the file for reading
	file, err := os.OpenFile(path, os.O_RDONLY, storePermissions)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Decode the YAML content from the file
	var store *ProjectStore
	if err := yaml.NewDecoder(file).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}

// readGlobalFile reads a global file
func (s *BaseStore) readGlobalFile(path string) (*GlobalStore, error) {
	// Check if the file exists
	if !util.Exists(path) {
		return nil, fmt.Errorf("store not found at %s", path)
	}
	// Open the file for reading
	file, err := os.OpenFile(path, os.O_RDONLY, storePermissions)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Decode the YAML content from the file
	var store *GlobalStore
	if err := yaml.NewDecoder(file).Decode(&store); err != nil {
		return nil, err
	}
	return store, nil
}

// updateFile updates a file
func (s *BaseStore) updateFile(path string, content interface{}) error {
	// Check if the file exists
	if !util.Exists(path) {
		return fmt.Errorf("store not found at %s", path)
	}
	// Open the file for writing
	file, err := os.OpenFile(path, os.O_WRONLY, storePermissions)
	if err != nil {
		return err
	}
	defer file.Close()
	// Encode the content to YAML and write it to the file
	if err := yaml.NewEncoder(file).Encode(content); err != nil {
		return err
	}
	return nil
}
