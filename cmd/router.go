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
	"regexp"

	"github.com/MohamedBeydoun/atlas/pkg/generator"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router [flags] [name]",
	Short: "Generates an express router.",
	Long:  `Generates a new express router with it's respective controller.`,
	RunE:  generateRouter,
}

func init() {
	generateCmd.AddCommand(routerCmd)
}

func generateRouter(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("router name not provided")
	} else if len(args) > 1 {
		return errors.New("too many arguments")
	}

	name := strcase.ToLowerCamel(args[0])

	validName, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9]*$`, name)
	if err != nil {
		return err
	}
	if !validName {
		return errors.New("invalid router name")
	}

	router, err := generator.NewRouter(name)
	if err != nil {
		return err
	}

	err = router.Create()
	return err
}
