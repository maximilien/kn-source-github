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
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/maximilien/kn-source-github/pkg/types"

	"github.com/spf13/pflag"
)

type gitHubSourceFlagsFactory struct {
	defaultFlagsFactory sourcetypes.FlagsFactory
	gitHubSourceFactory types.GitHubSourceFactory
}

func NewGitHubSourceFlagsFactory(gitHubSourceFactory types.GitHubSourceFactory) types.GitHubFlagsFactory {
	return &gitHubSourceFlagsFactory{
		defaultFlagsFactory: sourcefactories.NewDefaultFlagsFactory(gitHubSourceFactory),
		gitHubSourceFactory: gitHubSourceFactory,
	}
}

func (f *gitHubSourceFlagsFactory) KnSourceFactory() sourcetypes.KnSourceFactory {
	return f.gitHubSourceFactory
}

func (f *gitHubSourceFlagsFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.gitHubSourceFactory.KnSourceParams()
}

func (f *gitHubSourceFlagsFactory) GitHubSourceParams() *types.GitHubSourceParams {
	return f.gitHubSourceFactory.GitHubSourceParams()
}

func (f *gitHubSourceFlagsFactory) GitHubSourceFactory() types.GitHubSourceFactory {
	return f.gitHubSourceFactory
}

func (f *gitHubSourceFlagsFactory) CreateFlags() *pflag.FlagSet {
	flagSet := f.defaultFlagsFactory.CreateFlags()
	//TODO: add GitHub source flags
	return flagSet
}

func (f *gitHubSourceFlagsFactory) DeleteFlags() *pflag.FlagSet {
	flagSet := f.defaultFlagsFactory.DeleteFlags()
	//TODO: add GitHub source flags
	return flagSet
}

func (f *gitHubSourceFlagsFactory) UpdateFlags() *pflag.FlagSet {
	flagSet := f.defaultFlagsFactory.UpdateFlags()
	//TODO: add GitHub source flags
	return flagSet
}

func (f *gitHubSourceFlagsFactory) DescribeFlags() *pflag.FlagSet {
	flagSet := f.defaultFlagsFactory.DescribeFlags()
	//TODO: add GitHub source flags
	return flagSet
}
