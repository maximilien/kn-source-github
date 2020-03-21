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

	"github.com/maximilien/kn-source-github/pkg/types"

	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type gitHubSourceRunEFactory struct {
	gitHubSourceParams *types.GitHubSourceParams
	gitHubSourceClient types.GitHubSourceClient
}

func NewGitHubSourceRunEFactory(gitHubSourceParams *types.GitHubSourceParams,
	gitHubSourceClientFactory types.GitHubSourceClientFactory) types.GitHubSourceRunEFactory {
	return &gitHubSourceRunEFactory{
		gitHubSourceParams: gitHubSourceParams,
		gitHubSourceClient: gitHubSourceClientFactory.CreateGitHubSourceClient(),
	}
}

func (f *gitHubSourceRunEFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.GitHubSourceParams().KnSourceParams
}

func (f *gitHubSourceRunEFactory) GitHubSourceParams() *types.GitHubSourceParams {
	return f.gitHubSourceParams
}

func (f *gitHubSourceRunEFactory) KnSourceClient() sourcetypes.KnSourceClient {
	return f.GitHubSourceClient()
}

func (f *gitHubSourceRunEFactory) GitHubSourceClient() types.GitHubSourceClient {
	return f.gitHubSourceClient
}

func (f *gitHubSourceRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.GitHubSourceClient())
		return nil
	}
}

func (f *gitHubSourceRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.GitHubSourceClient())
		return nil
	}
}

func (f *gitHubSourceRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.GitHubSourceClient())
		return nil
	}
}

func (f *gitHubSourceRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.GitHubSourceClient())
		return nil
	}
}
