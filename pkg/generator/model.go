package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
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
	fmt.Printf(emoji.Sprintf(":gear:")+" Generating resources for the "+aurora.Yellow("%s").String()+" model\n\n", m.Name)

	overwriteModel := true
	overwriteInterface := true
	overwriteInteractions := true

	// check if model exists and if user wants to overwrite it
	if _, err := os.Stat(fmt.Sprintf("%s/src/database/models/%s.ts", m.Project.AbsolutePath, m.Name)); err == nil {
		overwriteModel = util.AskForConfirmation(fmt.Sprintf(aurora.Yellow("    src/database/models/%s.ts already exists. Would you like to overwrite it?").String(), m.Name))
	}
	if overwriteModel {
		fmt.Print("    src/database/models/")
		err := util.CreateFile(m, m.Name+".ts", m.Project.AbsolutePath+"/src/database/models", string(tpl.ModelTemplate()), 0)
		if err != nil {
			return err
		}
	}

	// check if interface exists and if user wants to overwrite it
	if _, err := os.Stat(fmt.Sprintf("%s/src/interfaces/%s.ts", m.Project.AbsolutePath, "I"+strings.Title(m.Name))); err == nil {
		overwriteInterface = util.AskForConfirmation(fmt.Sprintf(aurora.Yellow("    src/interfaces/I%s.ts already exists. Would you like to overwrite it?").String(), strings.Title(m.Name)))
	}
	if overwriteInterface {
		fmt.Print("    src/interfaces/")
		err := util.CreateFile(m, "I"+strings.Title(m.Name)+".ts", m.Project.AbsolutePath+"/src/interfaces", string(tpl.InterfaceTemplate()), 0)
		if err != nil {
			return err
		}
	}

	// check if interactions exists and if user wants to overwrite them
	if _, err := os.Stat(fmt.Sprintf("%s/src/database/interactions/%s.ts", m.Project.AbsolutePath, m.Name)); err == nil {
		overwriteInteractions = util.AskForConfirmation(fmt.Sprintf(aurora.Yellow("    src/database/interactions/%s.ts already exists. Would you like to overwrite it?").String(), m.Name))
	}
	if overwriteInteractions {
		fmt.Print("    src/database/interactions/")
		err := util.CreateFile(m, m.Name+".ts", m.Project.AbsolutePath+"/src/database/interactions", string(tpl.InteractionsTemplate()), 0)
		if err != nil {
			return err
		}
	}

	fmt.Println("\n" + emoji.Sprintf(":party_popper:") + "Done")

	return nil
}
