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

	"github.com/spf13/pflag"
)

type GitHubFlagsFactory struct{}

func NewGitHubFlagsFactory() commands.FlagsFactory {
	return &GitHubFlagsFactory{}
}

func (f *GitHubFlagsFactory) CreateFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("create", pflag.ExitOnError)
}

func (f *GitHubFlagsFactory) DeleteFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("delete", pflag.ExitOnError)
}

func (f *GitHubFlagsFactory) UpdateFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("create", pflag.ExitOnError)
}

func (f *GitHubFlagsFactory) DescribeFlags() *pflag.FlagSet {
	return pflag.NewFlagSet("create", pflag.ExitOnError)
}
