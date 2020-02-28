package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/MohamedBeydoun/atlas/pkg/prj"
	"github.com/MohamedBeydoun/atlas/pkg/util"
)

func getTestRouter() *Router {
	return &Router{
		Name:         "test",
		AbsolutePath: "/tmp/test/src/routes",
		Routes:       map[string]string{"post": "/tests"},
		Project: &prj.Project{
			Name:         "test",
			AbsolutePath: "/tmp/test",
			Port:         "3000",
			DBURL:        "mongodb://localhost:27017/PROJECT_NAME",
		},
	}
}

func TestGenerateRouter(t *testing.T) {
	router := getTestRouter()
	if err := router.Project.Create(); err != nil {
		t.Fatal(err)
	}
	if err := router.Create(); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(router.Project.AbsolutePath)

	expectedFiles := []string{"src/routes/test.ts"}

	for _, file := range expectedFiles {
		generatedFile := fmt.Sprintf("%s/%s", router.Project.AbsolutePath, file)
		goldenFile := fmt.Sprintf("../testdata/%s.golden", file)
		err := util.CompareFiles(generatedFile, goldenFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}
