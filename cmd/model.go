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
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model [flags] [arg]",
	Short: "Model generates a mongodb model.",
	Long: `Model generates a new mongodb model with the given
fields.`,
	RunE: generateModel,
}

func init() {
	generateCmd.AddCommand(modelCmd)

	modelCmd.Flags().StringToStringP("fields", "f", map[string]string{}, "Specify a field name and type")
	modelCmd.MarkFlagRequired("fields")
}

func generateModel(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Model name not provided\n")
	} else if len(args) > 1 {
		return errors.New("Too many arguments\n")
	}

	fields, err := cmd.Flags().GetStringToString("fields")
	if err != nil {
		return errors.New(err.Error())
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	model, err := generator.NewModel(args[0], fields, wd+"/src/database")
	if err != nil {
		return err
	}

	err = model.Create()
	return err
}
