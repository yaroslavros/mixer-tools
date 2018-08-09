// Copyright © 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/clearlinux/mixer-tools/builder"

	"github.com/spf13/cobra"
)

// Top level repo command ('mixer repo')
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Add, list, or remove RPM repositories for use by mixer",
}

var addRepoCmd = &cobra.Command{
	Use:   "add <name> <url>",
	Short: "Add repo <name> at <url>",
	Long:  `Add the repo at <url> as a repo from which to pull RPMs for building bundles`,
	Args:  cobra.ExactArgs(2),
	Run:   runAddRepo,
}

var removeRepoCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Removes repo <name> from the DNF conf used by mixer",
	Long:  `Remove the repo named <name> from the configured DNF conf used by mixer`,
	Args:  cobra.ExactArgs(1),
	Run:   runRemoveRepo,
}

var listReposCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured RPM Repositories",
	Long:  `List all RPM repositories configured in the DNF configuration file used by mixer`,
	Run:   runListRepos,
}

var initRepoCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the DNF conf with default repo enabled",
	Long:  `Initialize the DNF configuration file with the default 'Clear' repo enabled`,
	Run:   runInitRepo,
}

var setURLRepoCmd = &cobra.Command{
	Use:   "set-url <name> <url>",
	Short: "Sets the URL for repo <name> to <url>",
	Long:  `Sets the URL, for repo <name> to <url>. If <name> does not exist the repo will be added to the configuration.`,
	Args:  cobra.ExactArgs(2),
	Run:   runSetURLRepo,
}

var setExcludesRepoCmd = &cobra.Command{
	Use:   "exclude <repo> <pkg> [<pkg>...]",
	Short: "Exclude packages from a specified repo",
	Long:  `Exclude packages from a specified repo. These packages will be ignored during build bundles. Globbing is supported.`,
	Args:  cobra.MinimumNArgs(2),
	Run:   runExcludesRepo,
}

var repoCmds = []*cobra.Command{
	addRepoCmd,
	removeRepoCmd,
	listReposCmd,
	initRepoCmd,
	setURLRepoCmd,
	setExcludesRepoCmd,
}

func init() {
	for _, cmd := range repoCmds {
		repoCmd.AddCommand(cmd)
		cmd.Flags().StringVarP(&configFile, "config", "c", "", "Builder config to use")
	}

	RootCmd.AddCommand(repoCmd)
}

func runExcludesRepo(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.SetExcludesRepo(args[0], strings.Join(args[1:], " "))
	if err != nil {
		fail(err)
	}
	fmt.Printf("Excluded packages from repo %s:\n%s\n", args[0], strings.Join(args[1:], "\n"))
}

func runAddRepo(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.AddRepo(args[0], args[1])
	if err != nil {
		fail(err)
	}
	fmt.Printf("Added %s repo at %s url.\n", args[0], args[1])
}

func runRemoveRepo(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.RemoveRepo(args[0])
	if err != nil {
		fail(err)
	}
	fmt.Printf("Removed %s repo.\n", args[0])
}

func runListRepos(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.ListRepos()
	if err != nil {
		fail(err)
	}
}

func runInitRepo(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.NewDNFConfIfNeeded()
	if err != nil {
		fail(err)
	}
}

func runSetURLRepo(cmd *cobra.Command, args []string) {
	b, err := builder.NewFromConfig(configFile)
	if err != nil {
		fail(err)
	}

	err = b.SetURLRepo(args[0], args[1])
	if err != nil {
		fail(err)
	}
	fmt.Printf("Set %s baseurl to %s.\n", args[0], args[1])
}
