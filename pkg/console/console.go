package console

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/prj"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

type consoleConfig struct {
	DBURL  string
	Models []string
}

// Run starts a development console
func Run(dbURL string) error {
	fmt.Println("Running console...")

	project, err := prj.Current()
	if err != nil {
		return err
	}
	projectName := filepath.Base(project.AbsolutePath)
	dbURL = strings.Replace(dbURL, "PROJECT_NAME", projectName, -1)

	err = exec.Command("npm", "run", "build").Run()
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return errors.New(err.Error())
	}
	models := []string{}
	err = filepath.Walk("src/database/models/", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".") && info.Name() != "" {
			models = append(models, strings.Split(info.Name(), ".")[0])
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Print("Created ")
	err = util.CreateFile(consoleConfig{DBURL: dbURL, Models: models}, ".console", wd, string(tpl.ConsoleTemplate()), 0)
	if err != nil {
		return err
	}

	console := exec.Command("node", "--experimental-repl-await", wd+"/.console")
	console.Stdout = os.Stdout
	console.Stdin = os.Stdin
	console.Stderr = os.Stderr
	console.Run()

	return nil
}
