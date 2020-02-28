package generator

import (
	"fmt"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

// Controller holds the controller information
type Controller struct {
	Name         string
	Functions    []string
	AbsolutePath string
	Project      *prj.Project
}

// NewController creates a new controller struct
func NewController(name string, functions []string, path string) (*Controller, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Controller{
		Name:         name,
		Functions:    functions,
		AbsolutePath: path,
		Project:      project,
	}, nil
}

// Create generates the controller files
func (c *Controller) Create() error {
	fmt.Printf("Creating %s controller\n", c.Name)

	fmt.Printf("    %s/src/controllers/", c.Project.Name)
	err := util.CreateFile(c, c.Name+".ts", c.AbsolutePath, string(tpl.ControllerTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Println("Done")
	return nil
}
