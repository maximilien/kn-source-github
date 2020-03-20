// Copyright © 2018 The Knative Authors
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

package factories

import (
	sourcefactories "github.com/maximilien/kn-source-pkg/pkg/factories"

	"github.com/maximilien/kn-source-pkg/pkg/commands"

	"github.com/spf13/cobra"
)

type GitHubCommandFactory struct {
	defaultFactory commands.CommandFactory
}

func NewGitHubCommandFactory() commands.CommandFactory {
	return &GitHubCommandFactory{
		defaultFactory: sourcefactories.NewDefaultCommandFactory(),
	}
}

func (f *GitHubCommandFactory) SourceCommand() *cobra.Command {
	sourceCmd := f.defaultFactory.SourceCommand()
	sourceCmd.Use = "github"
	sourceCmd.Short = "Knative eventing GitHub source plugin"
	sourceCmd.Long = "Manage your Knative GitHub eventing sources"
	return sourceCmd
}

func (f *GitHubCommandFactory) CreateCommand(params *commands.KnSourceParams) *cobra.Command {
	createCmd := f.defaultFactory.CreateCommand(params)
	createCmd.Short = "create NAME"
	return createCmd
}

func (f *GitHubCommandFactory) DeleteCommand(params *commands.KnSourceParams) *cobra.Command {
	deleteCmd := f.defaultFactory.DeleteCommand(params)
	deleteCmd.Short = "delete NAME"
	deleteCmd.Long = "delete a GitHub source"
	return deleteCmd
}

func (f *GitHubCommandFactory) UpdateCommand(params *commands.KnSourceParams) *cobra.Command {
	updateCmd := f.defaultFactory.CreateCommand(params)
	updateCmd.Short = "update NAME"
	updateCmd.Long = "update a GitHub source"
	return updateCmd
}

func (f *GitHubCommandFactory) DescribeCommand(params *commands.KnSourceParams) *cobra.Command {
	describeCmd := f.defaultFactory.CreateCommand(params)
	describeCmd.Short = "describe NAME"
	describeCmd.Long = "update a GitHub source"
	return describeCmd
}
