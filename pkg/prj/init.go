package prj

import (
	"fmt"
	"html/template"
	"os"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
)

type Project struct {
	Name         string
	AbsolutePath string
	Port         string
	DBURL        string
}

// Create creates a new project directory with with a conventional express-typescript file structure
func (p *Project) Create() error {
	// create root directory
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		fmt.Printf("Creating new application \"%v\" at %v\n", p.Name, p.AbsolutePath)
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// create folders
	folders := []string{"src", "test"}
	for _, folder := range folders {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, folder)); os.IsNotExist(err) {
			if err := os.Mkdir(fmt.Sprintf("%s/%s", p.AbsolutePath, folder), 0751); err != nil {
				return err
			}
		}
	}

	// populate new folders
	err := p.populateSrc()
	if err != nil {
		return err
	}
	err = p.populateTest()
	if err != nil {
		return err
	}

	// create README.md
	readmeFile, err := os.Create(fmt.Sprintf("%s/README.md", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	readmeTemplate := template.Must(template.New("readme").Parse(string(tpl.ReadmeTemplate())))
	err = readmeTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	fmt.Println("\tcreated README.md")

	// create .gitignore
	gitignoreFile, err := os.Create(fmt.Sprintf("%s/.gitignore", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()

	gitignoreTemplate := template.Must(template.New("gitignore").Parse(string(tpl.GitignoreTemplate())))
	err = gitignoreTemplate.Execute(gitignoreFile, p)
	if err != nil {
		return err
	}

	fmt.Println("\tcreated .gitignore")

	fmt.Println("Done")
	return nil
}

func (p *Project) populateSrc() error {
	return nil
}

func (p *Project) populateTest() error {
	return nil
}
