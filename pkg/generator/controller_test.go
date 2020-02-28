package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

func getTestController() *Controller {
	return &Controller{
		Name:         "test",
		AbsolutePath: "/tmp/test/src/controllers",
		Functions:    []string{"index", "create"},
		Project: &prj.Project{
			Name:         "test",
			AbsolutePath: "/tmp/test",
			Port:         "3000",
			DBURL:        "mongodb://localhost:27017/PROJECT_NAME",
		},
	}
}

func TestGenerateController(t *testing.T) {
	controller := getTestController()
	if err := controller.Project.Create(); err != nil {
		t.Fatal(err)
	}
	if err := controller.Create(); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(controller.Project.AbsolutePath)

	expectedFiles := []string{"src/controllers/test.ts"}

	for _, file := range expectedFiles {
		generatedFile := fmt.Sprintf("%s/%s", controller.Project.AbsolutePath, file)
		goldenFile := fmt.Sprintf("../testdata/%s.golden", file)
		err := util.CompareFiles(generatedFile, goldenFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}
