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

// import (
// 	"errors"
// 	"os"

// 	"github.com/MohamedBeydoun/atlas/pkg/generator"
// 	"github.com/iancoleman/strcase"
// 	"github.com/spf13/cobra"
// )

// // controllerCmd represents the controller command
// var controllerCmd = &cobra.Command{
// 	Use:   "controller [flags] [name]",
// 	Short: "Controller generates an express controller.",
// 	Long: `Controller generates a new express controller with the given
// name and suggested functions.`,
// 	RunE: generateController,
// }

// func init() {
// 	generateCmd.AddCommand(controllerCmd)
// 	controllerCmd.Flags().StringSliceP("functions", "f", []string{}, "Specify functions e.g. index,show,create")
// }

// func generateController(cmd *cobra.Command, args []string) error {
// 	if len(args) == 0 {
// 		return errors.New("Controller name not provided\n")
// 	} else if len(args) > 1 {
// 		return errors.New("Too many arguments\n")
// 	}

// 	name := strcase.ToLowerCamel(args[0])
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}
// 	functions, err := cmd.Flags().GetStringSlice("functions")
// 	if err != nil {
// 		return errors.New(err.Error())
// 	}

// 	for i, _ := range functions {
// 		functions[i] = strcase.ToLowerCamel(string(functions[i]))
// 	}

// 	controller, err := generator.NewController(name, functions, wd+"/src/controllers")
// 	if err != nil {
// 		return err
// 	}

// 	err = controller.Create()
// 	return err
// }
