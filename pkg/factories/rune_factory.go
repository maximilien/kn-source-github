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
	gitHubSourceClient  types.GitHubSourceClient
}

func NewGitHubSourceRunEFactory(gitHubSourceFactory types.GitHubSourceFactory) types.GitHubRunEFactory {
	return &gitHubRunEFactory{
		gitHubSourceFactory: gitHubSourceFactory,
		gitHubSourceClient:  gitHubSourceFactory.GitHubSourceClient(),
	}
}

func (f *gitHubRunEFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.GitHubSourceFactory().KnSourceParams()
}

func (f *gitHubRunEFactory) KnSourceClient(cmd *cobra.Command) (sourcetypes.KnSourceClient, error) {
	return f.GitHubSourceFactory().GitHubSourceClient(), nil
}

func (f *gitHubRunEFactory) GitHubSourceClient(cmd *cobra.Command) (types.GitHubSourceClient, error) {
	knParams := f.gitHubSourceFactory.KnSourceParams().KnParams
	namespace, err := knParams.GetNamespace(cmd)
	if err != nil {
		return nil, err
	}

	f.gitHubSourceClient = f.gitHubSourceFactory.CreateKnSourceClient(namespace)

	return f.gitHubSourceClient, nil
}

func (f *gitHubRunEFactory) KnSourceFactory() sourcetypes.KnSourceFactory {
	return f.gitHubSourceFactory
}

func (f *gitHubRunEFactory) GitHubSourceFactory() types.GitHubSourceFactory {
	return f.gitHubSourceFactory
}

func (f *gitHubRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		f.gitHubSourceClient, err = f.GitHubSourceClient(cmd)
		if err != nil {
			return fmt.Errorf("could not access gitHubSourceClient for command %s", cmd.Name())
		}

		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		f.gitHubSourceClient, err = f.GitHubSourceClient(cmd)
		if err != nil {
			return fmt.Errorf("could not access gitHubSourceClient for command %s", cmd.Name())
		}

		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		f.gitHubSourceClient, err = f.GitHubSourceClient(cmd)
		if err != nil {
			return fmt.Errorf("could not access gitHubSourceClient for command %s", cmd.Name())
		}

		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.gitHubSourceClient)
		return nil
	}
}

func (f *gitHubRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		f.gitHubSourceClient, err = f.GitHubSourceClient(cmd)
		if err != nil {
			return fmt.Errorf("could not access gitHubSourceClient for command %s", cmd.Name())
		}

		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v\n", cmd.Name(), args, f.gitHubSourceClient)
		return nil
	}
}
