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
	"github.com/kyokomi/emoji"

	"github.com/MohamedBeydoun/atlas/pkg/tpl"
	"github.com/MohamedBeydoun/atlas/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dependencies = map[string]string{
		"mongoose": "/node_modules/mongoose",
		"express":  "/node_modules/express",
	}
)

type consoleConfig struct {
	DBURL  string
	Models []string
}

// Run starts a development console
func Run(dbURL string) error {
	fmt.Println(emoji.Sprintf(":red_circle:") + "Running console...\n")

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
		return errors.New("console must be run in the project's root directory")
	}

	// check deps
	if _, err := os.Stat(project.AbsolutePath + "/node_modules"); os.IsNotExist(err) {
		return errors.New("must run \"npm install\" first")
	}

	for depName, relativePath := range dependencies {
		if _, err := os.Stat(project.AbsolutePath + relativePath); os.IsNotExist(err) {
			return fmt.Errorf("missing package \"%s\"", depName)
		}
	}

	// check mongo
	mongoURI := strings.Split(dbURL, "/")
	mongoURI = mongoURI[:len(mongoURI)-1]
	longCtx, cancelLong := context.WithTimeout(context.Background(), 10*time.Second)
	shortCtx, cancelShort := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancelLong()
	defer cancelShort()

	fmt.Println(emoji.Sprintf(":floppy_disk:") + "Connecting to MongoDB...\n")
	client, err := mongo.Connect(longCtx, options.Client().ApplyURI(strings.Join(mongoURI, "/")))
	if err != nil {
		return fmt.Errorf("could not connect to mongo on %s", dbURL)
	}
	err = client.Ping(shortCtx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("could not connect to mongo on %s", dbURL)
	}
	client.Disconnect(shortCtx)

	fmt.Println(emoji.Sprintf(":hammer:") + "Building project...\n")
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

	fmt.Print(emoji.Sprintf(":memo:") + "Created ")
	err = util.CreateFile(consoleConfig{DBURL: dbURL, Models: models}, ".console", wd, string(tpl.ConsoleTemplate()), 0)
	if err != nil {
		return err
	}

	fmt.Println()
	console := exec.Command("node", "--experimental-repl-await", wd+"/.console")
	console.Stdout = os.Stdout
	console.Stdin = os.Stdin
	console.Stderr = os.Stderr
	console.Run()

	return nil
}
