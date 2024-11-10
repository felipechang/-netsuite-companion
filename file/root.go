package file

import (
	"log"
	"os"
	"path/filepath"
)

// Tree represents a file tree structure
type Tree struct {
	dirname string
}

// CreateTree creates a new Tree instance
func CreateTree() *Tree {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return &Tree{dirname}
}

// Build builds the file tree structure
func (s *Tree) Build() error {
	err := os.MkdirAll(filepath.Join(s.dirname, "src", "AccountConfiguration"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "SuiteScripts"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Templates"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Templates", "E-mail Templates"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Templates", "Marketing Templates"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Web Site Hosting Files"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Web Site Hosting Files", "Live Hosting Files"), os.ModePerm)
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "FileCabinet", "Web Site Hosting Files", "Staging Hosting Files"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "Objects"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(s.dirname, "src", "Translations"), os.ModePerm)
	if err != nil {
		return err
	}

	err = s.createFile(filepath.Join(s.dirname, "src", "deploy.xml"), `<deploy>
    <configuration>
        <path>~/AccountConfiguration/*</path>
    </configuration>
    <files>
        <path>~/FileCabinet/*</path>
    </files>
    <objects>
        <path>~/Objects/*</path>
    </objects>
    <translationimports>
        <path>~/Translations/*</path>
    </translationimports>
</deploy>`)
	if err != nil {
		return err
	}
	err = s.createFile(filepath.Join(s.dirname, ".gitignore"), `.idea
node_modules
	`)
	if err != nil {
		return err
	}
	err = s.createFile(filepath.Join(s.dirname, "package.json"), `{
  "name": "my-project",
  "version": "1.0.0",
  "devDependencies": {
    "@hitc/netsuite-types": "^2024.2.2",
    "@types/node": "^22.7.5",
    "typescript": "^5.6.2"
  }
}`)
	if err != nil {
		return err
	}
	err = s.createFile(filepath.Join(s.dirname, "tsconfig.json"), `{
  "compilerOptions": {
    "target": "es5",
    "module": "umd",
    "moduleResolution": "node",
    "sourceMap": false,
    "newLine": "LF",
    "experimentalDecorators": true,
    "noImplicitAny": true,
    "noImplicitThis": true,
    "strictNullChecks": true,
    "strictFunctionTypes": true,
    "strictPropertyInitialization": true,
    "baseUrl": "./",
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "lib": [
      "es5",
      "es2015.promise",
      "dom"
    ],
    "paths": {
      "N": [
        "node_modules/@hitc/netsuite-types/N"
      ],
      "N/*": [
        "node_modules/@hitc/netsuite-types/N/*"
      ]
    }
  },
  "include": [
    "src"
  ]
}`)
	if err != nil {
		return err
	}
	return nil
}

// createFile creates a new file
func (s *Tree) createFile(destination string, content string) error {
	gitignoreContent := []byte(content)
	err := os.WriteFile(destination, gitignoreContent, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
