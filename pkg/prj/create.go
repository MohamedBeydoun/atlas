package prj

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
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
	p.DBURL = strings.Replace(p.DBURL, "PROJECT_NAME", p.Name, -1)

	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		fmt.Printf("Creating new application \"%v\" at %v\n", p.Name, p.AbsolutePath)
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	} else {
		fmt.Printf("Application \"%s\" already exists at %v\n", p.Name, p.AbsolutePath)
		os.Exit(0)
	}

	err := util.CreateFolders(p, []string{"test"}, p.AbsolutePath, 0751, false, 1)
	if err != nil {
		return err
	}
	err = p.populateTest()
	if err != nil {
		return err
	}

	err = util.CreateFolders(p, []string{"src"}, p.AbsolutePath, 0751, false, 1)
	if err != nil {
		return err
	}
	err = p.populateSrc()
	if err != nil {
		return err
	}

	err = util.CreateFile(p, "package.json", p.AbsolutePath, string(tpl.PackageJSONTemplate()), 1)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, "tslint.json", p.AbsolutePath, string(tpl.TSLintTemplate()), 1)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, "tsconfig.json", p.AbsolutePath, string(tpl.TSConfigTemplate()), 1)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, "README.md", p.AbsolutePath, string(tpl.ReadmeTemplate()), 1)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, ".gitignore", p.AbsolutePath, string(tpl.GitignoreTemplate()), 1)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, ".atlas", p.AbsolutePath, "", 1)
	if err != nil {
		return err
	}

	fmt.Println("Done")

	fmt.Printf("\nRun the following commands:\n    cd %s\n    npm install\n", p.Name)
	return nil
}

// populateSrc populates the src directory with appropriate files and folders
func (p *Project) populateSrc() error {
	srcFolders := []string{"routes", "controllers", "interfaces", "middleware", "database"}
	err := util.CreateFolders(p, srcFolders, p.AbsolutePath+"/src", 0751, false, 2)
	if err != nil {
		return err
	}

	dbFolders := []string{"models", "interactions"}
	err = util.CreateFolders(p, dbFolders, p.AbsolutePath+"/src/database", 0751, false, 3)
	if err != nil {
		return err
	}

	utilFolder := []string{"util"}
	err = util.CreateFolders(p, utilFolder, p.AbsolutePath+"/src", 0751, false, 2)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, "statusCodes.ts", p.AbsolutePath+"/src/util", string(tpl.HTTPCodesTemplate()), 3)
	if err != nil {
		return err
	}

	err = util.CreateFile(p, "app.ts", p.AbsolutePath+"/src", string(tpl.AppTemplate()), 2)
	if err != nil {
		return err
	}
	err = util.CreateFile(p, "server.ts", p.AbsolutePath+"/src", string(tpl.ServerTemplate()), 2)
	if err != nil {
		return err
	}

	return nil
}

// populateSrc populates the test directory with appropriate files and folders
func (p *Project) populateTest() error {
	testFolders := []string{"routes"}
	err := util.CreateFolders(p, testFolders, p.AbsolutePath+"/test", 0751, false, 2)
	if err != nil {
		return err
	}

	return nil
}
