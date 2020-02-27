package generator

import (
	"fmt"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
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
	return nil
}
