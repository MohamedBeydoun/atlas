package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

// Router holds the router information
type Router struct {
	Name         string
	Routes       map[string]string
	AbsolutePath string
	Project      *prj.Project
}

// NewRouter creates a new router struct
func NewRouter(name string, path string) (*Router, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Router{
		Name:         name,
		AbsolutePath: path,
		Project:      project,
	}, nil
}

// Create generates the router files
func (r *Router) Create() error {
	fmt.Printf("Creating %s router\n", r.Name)

	err := error(nil)
	isNewRouter := true
	if _, err := os.Stat(fmt.Sprintf("%s/%s.ts", r.AbsolutePath, r.Name)); err == nil {
		proceed := util.AskForConfirmation(fmt.Sprintf("    src/routes/%s.ts already exists. Would you like to overwrite it?", r.Name))
		if !proceed {
			isNewRouter = false
			goto createController
		}
	}
	fmt.Print("    src/routes/")
	err = util.CreateFile(r, r.Name+".ts", r.AbsolutePath, string(tpl.RouterTemplate()), 0)
	if err != nil {
		return err
	}

createController:
	if _, err := os.Stat(fmt.Sprintf("%s/%s.ts", r.Project.AbsolutePath+"/src/controllers", r.Name)); err == nil {
		proceed := util.AskForConfirmation(fmt.Sprintf("    src/controllers/%s.ts already exists. Would you like to overwrite it?", r.Name))
		if !proceed {
			fmt.Println("Done")
			os.Exit(0)
		}
	}
	fmt.Print("    src/controllers/")
	err = util.CreateFile(r, r.Name+".ts", r.Project.AbsolutePath+"/src/controllers", string(tpl.ControllerTemplate()), 0)
	if err != nil {
		return err
	}

	// Update the app if new router
	if isNewRouter {
		fmt.Printf("    Updating src/app.ts\n")
		appFile, err := ioutil.ReadFile(fmt.Sprintf("%s/src/app.ts", r.Project.AbsolutePath))
		if err != nil {
			return err
		}
		appFileLines := strings.Split(string(appFile), "\n")
		useStr := fmt.Sprintf("app.use(%sRouter);", r.Name)
		importStr := fmt.Sprintf("import { %sRouter } from \"./routes/%s\";", r.Name, r.Name)
		linesToAdd := []string{useStr, importStr}
		for i, line := range appFileLines {
			if strings.Contains(line, "import cors from") {
				appFileLines = append(appFileLines, "")
				copy(appFileLines[i+2:], appFileLines[i+1:])
				appFileLines[i+1] = linesToAdd[1]
			}
			if strings.Contains(line, "app.use((req: Request") {
				appFileLines = append(appFileLines, "")
				copy(appFileLines[i+2:], appFileLines[i+1:])
				appFileLines[i+1] = linesToAdd[0]
			}
		}

		output := strings.Join(appFileLines, "\n")
		err = ioutil.WriteFile(fmt.Sprintf("%s/src/app.ts", r.Project.AbsolutePath), []byte(output), 0644)
		if err != nil {
			return err
		}
	}

	fmt.Println("Done")
	return nil
}
