package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

// Model holds the model information
type Model struct {
	Name    string
	Fields  map[string]string
	Project *prj.Project
}

// NewModel creates a new model struct
func NewModel(name string, fields map[string]string) (*Model, error) {
	project, err := prj.Current()
	if err != nil {
		return nil, err
	}

	return &Model{
		Name:    name,
		Fields:  fields,
		Project: project,
	}, nil
}

// Create generates the model files
func (m *Model) Create() error {
	fmt.Printf("Creating %s model\n", m.Name)
	err := error(nil)

	if _, err := os.Stat(fmt.Sprintf("%s/src/database/models/%s.ts", m.Project.AbsolutePath, m.Name)); err == nil {
		proceed := util.AskForConfirmation(fmt.Sprintf("    src/database/models/%s.ts already exists. Would you like to overwrite it?", m.Name))
		if !proceed {
			goto createInterface
		}
	}
	fmt.Print("    src/database/models/")
	err = util.CreateFile(m, m.Name+".ts", m.Project.AbsolutePath+"/src/database/models", string(tpl.ModelTemplate()), 0)
	if err != nil {
		return err
	}

createInterface:
	if _, err := os.Stat(fmt.Sprintf("%s/src/interfaces/%s.ts", m.Project.AbsolutePath, "I"+strings.Title(m.Name))); err == nil {
		proceed := util.AskForConfirmation(fmt.Sprintf("    src/interfaces/I%s.ts already exists. Would you like to overwrite it?", strings.Title(m.Name)))
		if !proceed {
			goto createInteractions
		}
	}
	fmt.Print("    src/interfaces/")
	err = util.CreateFile(m, "I"+strings.Title(m.Name)+".ts", m.Project.AbsolutePath+"/src/interfaces", string(tpl.InterfaceTemplate()), 0)
	if err != nil {
		return err
	}

createInteractions:
	if _, err := os.Stat(fmt.Sprintf("%s/src/database/interactions/%s.ts", m.Project.AbsolutePath, m.Name)); err == nil {
		proceed := util.AskForConfirmation(fmt.Sprintf("    src/database/interactions/%s.ts already exists. Would you like to overwrite it?", m.Name))
		if !proceed {
			fmt.Println("Done")
			os.Exit(0)
		}
	}
	fmt.Print("    src/database/interactions/")
	err = util.CreateFile(m, m.Name+".ts", m.Project.AbsolutePath+"/src/database/interactions", string(tpl.InteractionsTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Println("Done")

	return nil
}
