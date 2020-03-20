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
	"fmt"

	"github.com/maximilien/kn-source-pkg/pkg/commands"

	"github.com/spf13/cobra"
)

type GitHubRunEFactory struct{}

func NewGitHubRunEFactory() commands.RunEFactory {
	return &GitHubRunEFactory{}
}

func (f *GitHubRunEFactory) CreateRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("create RunE called for GitHub source: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *GitHubRunEFactory) DeleteRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("delete RunE called for GitHub source: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *GitHubRunEFactory) UpdateRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("update RunE called for GitHub source: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *GitHubRunEFactory) DescribeRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("describe RunE called for GitHub source: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}
