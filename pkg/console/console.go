package console

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

// Run starts a development console
func Run() error {
	fmt.Println("Running console...")

	models := []string{}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = filepath.Walk(wd+"/src/database/models/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".") {
			models = append(models, strings.Split(info.Name(), ".")[0])
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Print("Created ")
	err = util.CreateFile(models, "console.js", wd, string(tpl.ConsoleTemplate()), 0)
	if err != nil {
		return err
	}

	cmd := exec.Command("node", "-i", "-e", "\"$(< console.js)\"", "--experimental-repl-await")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
