package console

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/MohamedBeydoun/atlas/pkg/prj"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// check wd
	if wd != project.AbsolutePath {
		return errors.New("Console must be run in the project's root directory")
	}

	// check deps
	if _, err := os.Stat(project.AbsolutePath + "/node_modules"); os.IsNotExist(err) {
		return errors.New("Must run \"npm install\" first")
	}
	if _, err := os.Stat(project.AbsolutePath + "/node_modules/mongoose"); os.IsNotExist(err) {
		return errors.New("Missing package \"mongoose\"")
	}
	if _, err := os.Stat(project.AbsolutePath + "/node_modules/express"); os.IsNotExist(err) {
		return errors.New("Missing package \"express\"")
	}

	// check mongo
	mongoURI := strings.Split(dbURL, "/")
	mongoURI = mongoURI[:len(mongoURI)-1]
	longCtx, cancelLong := context.WithTimeout(context.Background(), 10*time.Second)
	shortCtx, cancelShort := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancelLong()
	defer cancelShort()

	client, err := mongo.Connect(longCtx, options.Client().ApplyURI(strings.Join(mongoURI, "/")))
	if err != nil {
		return errors.New(fmt.Sprintf("Could not connect to mongo on %s", dbURL))
	}
	err = client.Ping(shortCtx, readpref.Primary())
	if err != nil {
		return errors.New(fmt.Sprintf("Could not connect to mongo on %s", dbURL))
	}
	client.Disconnect(shortCtx)

	err = exec.Command("npm", "run", "build").Run()
	if err != nil {
		return err
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
