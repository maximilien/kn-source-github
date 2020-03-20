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
	"github.com/maximilien/kn-source-pkg/pkg/commands"
	"github.com/maximilien/kn-source-pkg/pkg/commands/source"

	"github.com/spf13/cobra"
)

type DefautCommandFactory struct{}

func NewDefaultCommandFactory() commands.CommandFactory {
	return &DefautCommandFactory{}
}

func (f *DefautCommandFactory) SourceCommand() *cobra.Command {
	return source.NewSourceCommand()
}

func (f *DefautCommandFactory) CreateCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewCreateCommand(params)
}

func (f *DefautCommandFactory) DeleteCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewDeleteCommand(params)
}

func (f *DefautCommandFactory) UpdateCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewUpdateCommand(params)
}

func (f *DefautCommandFactory) DescribeCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewDescribeCommand(params)
}
