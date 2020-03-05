package prj

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Current fetches the current project information
func Current() (*Project, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	counter := 0
	for i := 0; i < 4; i++ {
		if _, err := os.Stat(wd + "/.atlas"); os.IsNotExist(err) && i == 3 {
			return nil, errors.New("Not in an atlas project")
		}
		if _, err := os.Stat(wd + "/.atlas"); err == nil {
			break
		}

		counter++
		wd += "/.."
	}

	wdArr := strings.Split(wd, "/")
	for i := 0; i < counter*2; i++ {
		wdArr = wdArr[:len(wdArr)-1]
	}

	wd = strings.Join(wdArr, "/")

	project := Project{
		Name:         filepath.Base(wd),
		AbsolutePath: wd,
		Port:         "",
		DBURL:        "",
	}
	return &project, nil
}
