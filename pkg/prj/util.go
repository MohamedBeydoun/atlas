package prj

import (
	"errors"
	"os"
	"path/filepath"
)

// Current fetches the current project information
func Current() (*Project, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	project := Project{
		Name:         filepath.Base(wd),
		AbsolutePath: wd,
		Port:         "",
		DBURL:        "",
	}
	return &project, nil
}
