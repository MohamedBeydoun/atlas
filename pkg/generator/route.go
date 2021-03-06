package generator

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

var (
	newControllerBlock = `
    %s: async (req: Request, res: Response) => {
        try {
            res.status(statusCodes.NOT_IMPLEMENTED).send({ msg: "Not Implemented" });
        } catch (err) {
            res.status(statusCodes.INTERNAL_SERVER_ERROR).send(err);
        }
    },`
)

// Route holds the route information
type Route struct {
	Router     string
	Controller string
	URL        string
	Method     string
	Project    *prj.Project
}

// NewRoute creates a new route struct
func NewRoute(router string, method string, url string, controller string) (*Route, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Route{
		Router:     router,
		Controller: controller,
		URL:        url,
		Method:     method,
		Project:    project,
	}, nil
}

// Create generates the route files
func (r *Route) Create() error {
	fmt.Printf(emoji.Sprintf(":gear:")+" Generating route for the "+aurora.Yellow("%s").String()+" router\n\n", r.Router)

	// Update the router
	fmt.Printf("    "+aurora.Cyan("Updating ").String()+"src/routes/%s.ts\n", r.Router)
	routerFile, err := ioutil.ReadFile(fmt.Sprintf("%s/src/routes/%s.ts", r.Project.AbsolutePath, r.Router))
	if err != nil {
		return err
	}
	routerFileLines := strings.Split(string(routerFile), "\n")
	routerStr := fmt.Sprintf("\n%sRouter.%s(\"%s\", %sController.%s);", r.Router, r.Method, r.URL, r.Router, r.Controller)

	linesToAdd := []string{routerStr}
	for i, line := range routerFileLines {
		if strings.Contains(line, fmt.Sprintf("%sRouter: Router = Router()", r.Router)) {
			routerFileLines = append(routerFileLines, "")
			copy(routerFileLines[i+2:], routerFileLines[i+1:])
			routerFileLines[i+1] = linesToAdd[0]
		}
	}

	routerOutput := strings.Join(routerFileLines, "\n")
	err = ioutil.WriteFile(fmt.Sprintf("%s/src/routes/%s.ts", r.Project.AbsolutePath, r.Router), []byte(routerOutput), 0644)
	if err != nil {
		return err
	}

	// Update the controller
	fmt.Printf("    "+aurora.Cyan("Updating ").String()+"src/controllers/%s.ts\n", r.Router)
	controllerFile, err := ioutil.ReadFile(fmt.Sprintf("%s/src/controllers/%s.ts", r.Project.AbsolutePath, r.Router))
	if err != nil {
		return err
	}
	controllerFileLines := strings.Split(string(controllerFile), "\n")
	controllerStr := fmt.Sprintf(newControllerBlock, r.Controller)

	linesToAdd = []string{controllerStr}
	for i, line := range controllerFileLines {
		if strings.Contains(line, fmt.Sprintf("%sController = {", r.Router)) {
			controllerFileLines = append(controllerFileLines, "")
			copy(controllerFileLines[i+2:], controllerFileLines[i+1:])
			controllerFileLines[i+1] = linesToAdd[0]
		}
	}

	controllerOutput := strings.Join(controllerFileLines, "\n")
	err = ioutil.WriteFile(fmt.Sprintf("%s/src/controllers/%s.ts", r.Project.AbsolutePath, r.Router), []byte(controllerOutput), 0644)
	if err != nil {
		return err
	}

	fmt.Println("\n" + emoji.Sprintf(":party_popper:") + "Done")

	return nil
}
