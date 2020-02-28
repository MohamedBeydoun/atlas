package generator

import (
	"fmt"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

// Model holds the model information
type Model struct {
	Name         string
	Fields       map[string]string
	AbsolutePath string
	Project      *prj.Project
}

// NewModel creates a new model struct
func NewModel(name string, fields map[string]string, path string) (*Model, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Model{
		Name:         name,
		Fields:       fields,
		AbsolutePath: path,
		Project:      project,
	}, nil
}

// Create generates the model files
func (m *Model) Create() error {
	fmt.Printf("Creating %s model\n", m.Name)

	fmt.Printf("    %s/src/database/models/", m.Project.Name)
	err := util.CreateFile(m, m.Name+".ts", m.AbsolutePath+"/models", string(tpl.ModelTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Printf("    %s/src/interfaces/", m.Project.Name)
	err = util.CreateFile(m, "I"+strings.Title(m.Name)+".ts", m.Project.AbsolutePath+"/src/interfaces", string(tpl.InterfaceTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Printf("    %s/src/database/interactions/", m.Project.Name)
	err = util.CreateFile(m, m.Name+".ts", m.AbsolutePath+"/interactions", string(tpl.InteractionsTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Println("Done")

	return nil
}
