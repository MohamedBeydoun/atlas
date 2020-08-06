package prj

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

var (
	rootFolders = make([]folder, 2)
	rootFiles   = map[string]string{
		"package.json":  string(tpl.PackageJSONTemplate()),
		"tslint.json":   string(tpl.TSLintTemplate()),
		"tsconfig.json": string(tpl.TSConfigTemplate()),
		"README.md":     string(tpl.ReadmeTemplate()),
		".gitignore":    string(tpl.GitignoreTemplate()),
		".atlas":        "",
	}
)

// Project is the structure holding project information
type Project struct {
	Name         string
	AbsolutePath string
	Port         string
	DBURL        string
}

type folder struct {
	name       string
	permission os.FileMode
	path       string
	level      int
	populate   func() error
}

// NewProject initialized a new project's information
func NewProject(name, path, port, dbURL string) *Project {
	newProject := Project{
		Name:         name,
		AbsolutePath: path,
		Port:         port,
		DBURL:        dbURL,
	}

	rootFolders = []folder{
		{
			name:       "test",
			permission: 0751,
			level:      1,
			path:       path,
			populate:   newProject.populateTest,
		},
		{
			name:       "src",
			permission: 0751,
			level:      1,
			path:       path,
			populate:   newProject.populateSrc,
		},
	}

	return &newProject
}

// Create creates a new project directory with with a conventional express-typescript file structure
func (p *Project) Create() error {
	p.DBURL = strings.Replace(p.DBURL, "PROJECT_NAME", p.Name, -1)

	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		fmt.Printf(emoji.Sprint(":sparkles:")+"Creating project in"+aurora.Yellow(" %v\n\n").String(), p.AbsolutePath)
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	} else {
		fmt.Printf("Application \"%s\" already exists at %v\n", p.Name, p.AbsolutePath)
		os.Exit(0)
	}

	fmt.Printf(emoji.Sprint(":rocket:") + "Invoking generators...\n\n")
	for _, folder := range rootFolders {
		err := util.CreateFolders(p, []string{folder.name}, folder.path, folder.permission, false, folder.level)
		if err != nil {
			return err
		}

		err = folder.populate()
		if err != nil {
			return err
		}
	}

	for filename, template := range rootFiles {
		err := util.CreateFile(p, filename, p.AbsolutePath, template, 1)
		if err != nil {
			return err
		}
	}

	fmt.Printf("\n"+emoji.Sprint(":party_popper:")+"Successfully created project"+aurora.Yellow(" %v\n").String(), p.Name)
	fmt.Printf("\n"+emoji.Sprint(":point_right:")+"Get started with the following commands:\n\n"+aurora.Cyan("    $ cd %s\n    $ npm install\n").String(), p.Name)

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

// populateTest populates the test directory with appropriate files and folders
func (p *Project) populateTest() error {
	testFolders := []string{"routes"}
	err := util.CreateFolders(p, testFolders, p.AbsolutePath+"/test", 0751, false, 2)
	if err != nil {
		return err
	}

	return nil
}
