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

package types

import (
	"github.com/spf13/cobra"

	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type GitHubSourceClient interface {
	sourcetypes.KnSourceClient
}

type GitHubSourceFactory interface {
	sourcetypes.KnSourceFactory

	GitHubSourceParams() *GitHubSourceParams
	GitHubSourceClient() GitHubSourceClient

	CreateGitHubSourceParams() *GitHubSourceParams
	CreateGitHubSourceClient(namespace string) GitHubSourceClient
}

type GitHubCommandFactory interface {
	sourcetypes.CommandFactory

	GitHubSourceFactory() GitHubSourceFactory
}

type GitHubFlagsFactory interface {
	sourcetypes.FlagsFactory

	GitHubSourceFactory() GitHubSourceFactory
}

type GitHubRunEFactory interface {
	sourcetypes.RunEFactory

	GitHubSourceFactory() GitHubSourceFactory
	GitHubSourceClient(cmd *cobra.Command) (GitHubSourceClient, error)
}
