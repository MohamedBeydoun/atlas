/*
Copyright © 2020 Mohamed Beydoun

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

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [flags] [arg]",
	Short: "Create creates a new express project.",
	Long: `Create will create a new express project with given 
configs. The new project will follow the appropriate structure for an 
express-typescript project.`,
	RunE: createProject,
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().IntP("port", "p", 3000, "Specify port number")
	createCmd.Flags().String("db-url", "mongodb://localhost:27017/PROJECT_NAME", "Specify mongodb url")
}
func createProject(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("Project name not provided")
	} else if len(args) > 1 {
		return errors.New("Too many arguments")
	}

	fmt.Printf("Creating new project: %v\n", args[0])
	return nil
}
