package prj

import (
	"fmt"
	"html/template"
	"os"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
)

// Project is the structure holding project information
type Project struct {
	Name         string
	AbsolutePath string
	Port         string
	DBURL        string
}

// Create creates a new project directory with with a conventional express-typescript file structure
func (p *Project) Create() error {
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		fmt.Printf("Creating new application \"%v\" at %v\n", p.Name, p.AbsolutePath)
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	} else {
		fmt.Printf("Application \"%s\" already exists at %v\n", p.Name, p.AbsolutePath)
		os.Exit(0)
	}

	err := p.createFolders([]string{"test"}, p.AbsolutePath, 0751, true, 1)
	if err != nil {
		return err
	}
	err = p.populateTest()
	if err != nil {
		return err
	}

	err = p.createFolders([]string{"src"}, p.AbsolutePath, 0751, true, 1)
	if err != nil {
		return err
	}
	err = p.populateSrc()
	if err != nil {
		return err
	}

	err = p.createFile("package.json", p.AbsolutePath, string(tpl.PackageJSONTemplate()), 1)
	if err != nil {
		return err
	}
	err = p.createFile("tslint.json", p.AbsolutePath, string(tpl.TSLintTemplate()), 1)
	if err != nil {
		return err
	}
	err = p.createFile("tsconfig.json", p.AbsolutePath, string(tpl.TSConfigTemplate()), 1)
	if err != nil {
		return err
	}
	err = p.createFile("README.md", p.AbsolutePath, string(tpl.ReadmeTemplate()), 1)
	if err != nil {
		return err
	}
	err = p.createFile(".gitignore", p.AbsolutePath, string(tpl.GitignoreTemplate()), 1)
	if err != nil {
		return err
	}

	fmt.Println("Done")
	return nil
}

// populateSrc populates the src directory with appropriate files and folders
func (p *Project) populateSrc() error {
	srcFolders := []string{"routes", "controllers", "interfaces", "middleware", "util", "database"}
	err := p.createFolders(srcFolders, p.AbsolutePath+"/src", 0751, true, 2)
	if err != nil {
		return err
	}

	dbFolders := []string{"models", "interactions"}
	err = p.createFolders(dbFolders, p.AbsolutePath+"/src/database", 0751, true, 3)
	if err != nil {
		return err
	}

	err = p.createFile("app.ts", p.AbsolutePath+"/src", string(tpl.AppTemplate()), 2)
	if err != nil {
		return err
	}
	err = p.createFile("server.ts", p.AbsolutePath+"/src", string(tpl.ServerTemplate()), 2)
	if err != nil {
		return err
	}

	return nil
}

// populateSrc populates the test directory with appropriate files and folders
func (p *Project) populateTest() error {
	testFolders := []string{"routes"}
	err := p.createFolders(testFolders, p.AbsolutePath+"/test", 0751, true, 2)
	if err != nil {
		return err
	}

	return nil
}

// createFolders creates a list of folders
func (p *Project) createFolders(folders []string, path string, permissions os.FileMode, withKeep bool, level int) error {
	for _, folder := range folders {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", path, folder)); os.IsNotExist(err) {
			if err := os.Mkdir(fmt.Sprintf("%s/%s", path, folder), permissions); err != nil {
				return err
			}
			if withKeep {
				err := keepFolder(path + "/" + folder)
				if err != nil {
					return err
				}
			}

			for i := 0; i < level; i++ {
				fmt.Print("    ")
			}
			fmt.Printf("%s/\n", folder)
		}
	}

	return nil
}

// createFile creates a file
func (p *Project) createFile(name string, path string, templateString string, level int) error {
	file, err := os.Create(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	defer file.Close()

	fileTemplate := template.Must(template.New(name).Parse(templateString))
	err = fileTemplate.Execute(file, p)
	if err != nil {
		return err
	}

	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%s\n", name)

	return nil
}

// keepFolder creates a .keep file in directory to keep the folder
func keepFolder(path string) error {
	keepFile, err := os.Create(fmt.Sprintf("%s/.keep", path))
	if err != nil {
		return err
	}
	defer keepFile.Close()

	return nil
}
