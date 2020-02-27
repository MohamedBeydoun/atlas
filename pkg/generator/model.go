package generator

import (
	"fmt"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

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

func (m *Model) Create() error {
	fmt.Printf("Creating %s model under %s\n", m.Name, m.AbsolutePath)

	fmt.Println("    models/")
	err := util.CreateFile(m, m.Name+".ts", m.AbsolutePath+"/models", string(tpl.ModelTemplate()), 2)
	if err != nil {
		return err
	}

	fmt.Println("    interfaces/")
	err = util.CreateFile(m, "I"+strings.Title(m.Name)+".ts", m.Project.AbsolutePath+"/src/interfaces", string(tpl.InterfaceTemplate()), 2)
	if err != nil {
		return err
	}

	fmt.Println("    interactions/")
	err = util.CreateFile(m, m.Name+".ts", m.AbsolutePath+"/interactions", string(tpl.InteractionsTemplate()), 2)
	if err != nil {
		return err
	}
	return nil
}
