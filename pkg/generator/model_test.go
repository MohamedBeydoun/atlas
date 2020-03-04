package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

func getTestModel() *Model {
	return &Model{
		Name:   "test",
		Fields: map[string]string{"test": "string", "test2": "[string]"},
		Project: &prj.Project{
			Name:         "test",
			AbsolutePath: "/tmp/test",
			Port:         "3000",
			DBURL:        "mongodb://localhost:27017/PROJECT_NAME",
		},
	}
}

func TestGenerateModel(t *testing.T) {
	model := getTestModel()
	if err := model.Project.Create(); err != nil {
		t.Fatal(err)
	}
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(model.Project.AbsolutePath)

	expectedFiles := []string{"src/interfaces/ITest.ts", "src/database/interactions/test.ts", "src/database/models/test.ts"}

	for _, file := range expectedFiles {
		generatedFile := fmt.Sprintf("%s/%s", model.Project.AbsolutePath, file)
		goldenFile := fmt.Sprintf("../testdata/%s.golden", file)
		err := util.CompareFiles(generatedFile, goldenFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}
