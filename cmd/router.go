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
	"regexp"
	"strings"

	"github.com/MohamedBeydoun/atlas/pkg/generator"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router [flags] [name]",
	Short: "Router generates an express router",
	Long: `Router generates a new express router with the given
name and suggested functions.`,
	RunE: generateRouter,
}

func init() {
	generateCmd.AddCommand(routerCmd)
	routerCmd.Flags().StringToStringP("routes", "r", map[string]string{}, "Specify routes e.g. get=\"/users\",post=\"/users\"")
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
	routes := make(map[string]string)
	rawRoutes, err := cmd.Flags().GetStringToString("routes")
	if err != nil {
		return errors.New(err.Error())
	}

	expectedHttpMethods := []string{"get", "post", "put", "patch", "update", "delete"}
	for method, route := range rawRoutes {
		for _, expectedMethod := range expectedHttpMethods {
			if strings.ToLower(string(method)) == expectedMethod {
				break
			}
			if !(strings.ToLower(string(method)) == expectedMethod) && expectedMethod == "delete" {
				return errors.New(fmt.Sprintf("Unknown http method: %s\n", method))
			}
		}
		validURL, err := regexp.MatchString(`^\/[/.a-zA-Z0-9-]+$`, string(route))
		if err != nil {
			return err
		}
		if string(route[0]) != "/" || !validURL {
			return errors.New(fmt.Sprintf("Invalid routes format: %s\n", string(route)))
		}

		routes[strings.ToLower(method)] = route
	}

	router, err := generator.NewRouter(name, routes, wd+"/src/routes")
	if err != nil {
		return err
	}

	err = router.Create()
	return err
}
