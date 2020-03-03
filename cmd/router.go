/*
Copyright © 2020 Mohamed Beydoun mohamed.beydoun@mail.mcgill.ca

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"os"

	"github.com/MohamedBeydoun/atlas/pkg/generator"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router [flags] [name]",
	Short: "Router generates an express router.",
	Long:  `Router generates a new express router with it's respective controller.`,
	RunE:  generateRouter,
}

func init() {
	generateCmd.AddCommand(routerCmd)
}

func generateRouter(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Router name not provided\n")
	} else if len(args) > 1 {
		return errors.New("Too many arguments\n")
	}

	name := strcase.ToLowerCamel(args[0])
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	router, err := generator.NewRouter(name, wd+"/src/routes")
	if err != nil {
		return err
	}

	err = router.Create()
	return err
}
