package generator

import (
	"fmt"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
)

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

func (c *Controller) Create() error {
	fmt.Printf("Creating %s controller\n", c.Name)
	return nil
}
