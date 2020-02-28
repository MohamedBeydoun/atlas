package generator

import (
	"fmt"

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
func NewRouter(name string, routes map[string]string, path string) (*Router, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Router{
		Name:         name,
		Routes:       routes,
		AbsolutePath: path,
		Project:      project,
	}, nil
}

// Create generates the controller files
func (r *Router) Create() error {
	fmt.Printf("Creating %s router\n", r.Name)

	fmt.Printf("    %s/src/routes/", r.Project.Name)
	err := util.CreateFile(r, r.Name+".ts", r.AbsolutePath, string(tpl.RouterTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Println("Done")
	fmt.Println("\nDon't forget to use the router in your app.ts")
	return nil
}