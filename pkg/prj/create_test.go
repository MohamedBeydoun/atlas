package prj

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/MohamedBeydoun/atlas/pkg/util"
)

func getProject() *Project {
	return &Project{
		Name:         "test",
		AbsolutePath: "/tmp/test",
		Port:         "3000",
		DBURL:        "mongodb://localhost:27017/PROJECT_NAME",
	}
}

func TestCreateProject(t *testing.T) {
	project := getProject()
	if err := project.Create(); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(project.AbsolutePath)

	expectedFolders := []string{"test", "src", "src/routes", "src/controllers", "src/database", "src/interfaces", "src/util", "src/database/models", "src/database/interactions"}
	expectedFiles := []string{".gitignore", "README.md", "package.json", "tsconfig.json", "tslint.json", "src/app.ts", "src/server.ts"}

	for _, folder := range expectedFolders {
		_, err := os.Stat(fmt.Sprintf("%s/%s", project.AbsolutePath, folder))
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, file := range expectedFiles {
		generatedFile := fmt.Sprintf("%s/%s", project.AbsolutePath, file)
		goldenFile := fmt.Sprintf("../testdata/%s.golden", filepath.Base(file))
		err := util.CompareFiles(generatedFile, goldenFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}
