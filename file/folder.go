package file

import (
	"fmt"
	"netsuite-companion/store"
	"os"
	"path/filepath"
)

// CreateManifest creates a manifest file for a NetSuite project
func (s *Tree) CreateManifest(project *store.ProjectStore) error {
	err := s.createFile("src/manifest.xml", fmt.Sprintf(`<manifest projecttype="ACCOUNTCUSTOMIZATION">
  <projectname>{{%s}}</projectname>
  <frameworkversion>1.0</frameworkversion>
</manifest>
`, project.Current))
	if err != nil {
		return err
	}
	return nil
}

// CreateProjectFolder creates a project folder structure for a NetSuite project
func (s *Tree) CreateProjectFolder(global *store.GlobalStore, project *store.ProjectStore) error {
	err := os.MkdirAll(filepath.Join("src", "FileCabinet", "SuiteScripts", global.VendorName, project.Current), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
