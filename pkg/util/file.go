package util

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Resource interface{}

// CreateFolders creates a list of folders
func CreateFolders(r Resource, folders []string, path string, permissions os.FileMode, withKeep bool, level int) error {
	for _, folder := range folders {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", path, folder)); os.IsNotExist(err) {
			if err := os.Mkdir(fmt.Sprintf("%s/%s", path, folder), permissions); err != nil {
				return err
			}
			if withKeep {
				err := KeepFolder(path + "/" + folder)
				if err != nil {
					return err
				}
			}

			for i := 0; i < level; i++ {
				fmt.Print("    ")
			}
			fmt.Printf("%s/\n", folder)
		}
	}

	return nil
}

// CreateFile creates a file
func CreateFile(r Resource, name string, path string, templateString string, level int) error {
	tplFuncMap := template.FuncMap{
		"ToUpper":   strings.ToUpper,
		"ToTitle":   strings.Title,
		"TrimLeft":  func(str string) string { return str[1:] },
		"TrimRight": func(str string) string { n := len(str); return str[:n-1] },
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	defer file.Close()

	fileTemplate := template.Must(template.New(name).Funcs(tplFuncMap).Parse(templateString))
	err = fileTemplate.Execute(file, r)
	if err != nil {
		return err
	}

	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Printf("%s\n", name)

	return nil
}

// KeepFolder creates a .keep file in directory to keep the folder
func KeepFolder(path string) error {
	keepFile, err := os.Create(fmt.Sprintf("%s/.keep", path))
	if err != nil {
		return err
	}
	defer keepFile.Close()

	return nil
}
