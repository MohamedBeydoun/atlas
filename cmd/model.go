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
	"fmt"
	"os"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/generator"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model [flags] [name]",
	Short: "Model generates a mongodb model.",
	Long: `Model generates a new mongodb model with the given
fields.`,
	RunE: generateModel,
}

func init() {
	generateCmd.AddCommand(modelCmd)

	modelCmd.Flags().StringToStringP("fields", "f", map[string]string{}, "Specify field names and types e.g. name=string,friends=[string]")
	modelCmd.MarkFlagRequired("fields")
}

func generateModel(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Model name not provided\n")
	} else if len(args) > 1 {
		return errors.New("Too many arguments\n")
	}

	name := strcase.ToLowerCamel(args[0])
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fields := make(map[string]string)
	rawFields, err := cmd.Flags().GetStringToString("fields")
	if err != nil {
		return errors.New(err.Error())
	}

	allowedTypes := []string{"string", "boolean", "number", "symbol", "object", "[string]", "[boolean]", "[number]", "[symbol]"}
	for field, fieldType := range rawFields {
		for _, allowedType := range allowedTypes {
			if strings.ToLower(string(fieldType)) == allowedType {
				break
			}
			if !(strings.ToLower(string(fieldType)) == allowedType) && allowedType == "object" {
				return errors.New(fmt.Sprintf("Unknown type: %s\n", string(fieldType)))
			}
		}

		fields[strcase.ToLowerCamel(field)] = strings.ToLower(fieldType)
	}

	model, err := generator.NewModel(name, fields, wd)
	if err != nil {
		return err
	}

	err = model.Create()
	return err
}
