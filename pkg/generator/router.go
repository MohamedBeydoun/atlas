package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

// Router holds the router information
type Router struct {
	Name    string
	Project *prj.Project
}

// NewRouter creates a new router struct
func NewRouter(name string) (*Router, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Router{
		Name:    name,
		Project: project,
	}, nil
}

// Create generates the router files
func (r *Router) Create() error {
	fmt.Printf(emoji.Sprintf(":gear:")+" Generating "+aurora.Yellow("%s").String()+" router\n\n", r.Name)

	isNewRouter := true
	overwriteRouter := true
	overwriteController := true

	// check if router exists and if user wants to overwrite it
	if _, err := os.Stat(fmt.Sprintf("%s/src/routes/%s.ts", r.Project.AbsolutePath, r.Name)); err == nil {
		overwriteRouter = util.AskForConfirmation(fmt.Sprintf(aurora.Yellow("    src/routes/%s.ts already exists. Would you like to overwrite it?").String(), r.Name))
		isNewRouter = false
	}
	if overwriteRouter {
		fmt.Print("    src/routes/")
		err := util.CreateFile(r, r.Name+".ts", r.Project.AbsolutePath+"/src/routes", string(tpl.RouterTemplate()), 0)
		if err != nil {
			return err
		}
	}

	// check if controller exists and if user wants to overwrite it
	if _, err := os.Stat(fmt.Sprintf("%s/%s.ts", r.Project.AbsolutePath+"/src/controllers", r.Name)); err == nil {
		overwriteController = util.AskForConfirmation(fmt.Sprintf(aurora.Yellow("    src/controllers/%s.ts already exists. Would you like to overwrite it?").String(), r.Name))
	}
	if overwriteController {
		fmt.Print("    src/controllers/")
		err := util.CreateFile(r, r.Name+".ts", r.Project.AbsolutePath+"/src/controllers", string(tpl.ControllerTemplate()), 0)
		if err != nil {
			return err
		}
	}

	// Update the app if it's a new router
	if isNewRouter {
		fmt.Printf("    " + aurora.Cyan("Updating ").String() + "src/app.ts\n")
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

	fmt.Println("\n" + emoji.Sprintf(":party_popper:") + "Done")

	return nil
}
