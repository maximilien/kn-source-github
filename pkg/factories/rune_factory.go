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

type gitHubRunEFactory struct {
	gitHubSourceFactory types.GitHubSourceFactory
}

func NewGitHubSourceRunEFactory(gitHubSourceFactory types.GitHubSourceFactory) types.GitHubRunEFactory {
	return &gitHubRunEFactory{
		gitHubSourceFactory: gitHubSourceFactory,
	}
}

func (f *gitHubRunEFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.GitHubSourceFactory().KnSourceParams()
}

func (f *gitHubRunEFactory) KnSourceClient(namespace string) sourcetypes.KnSourceClient {
	return f.GitHubSourceFactory().CreateGitHubSourceClient(namespace)
}

func (f *gitHubRunEFactory) GitHubSourceClient(namespace string) types.GitHubSourceClient {
	return f.GitHubSourceFactory().CreateGitHubSourceClient(namespace)
}

func (f *gitHubRunEFactory) KnSourceFactory() sourcetypes.KnSourceFactory {
	return f.GitHubSourceFactory()
}

func (f *gitHubRunEFactory) GitHubSourceFactory() types.GitHubSourceFactory {
	return f.gitHubSourceFactory
}

// Factory *RunE methods

func (f *gitHubRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		gitHubSourceClient := f.GitHubSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		gitHubSourceClient := f.GitHubSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		gitHubSourceClient := f.GitHubSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		gitHubSourceClient := f.GitHubSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, gitHubSourceClient)
		return nil
	}
}
