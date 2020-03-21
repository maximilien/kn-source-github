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
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type GitHubSource interface {
	GitHubSourceParams() *GitHubSourceParams
}

type GitHubSourceClient interface {
	GitHubSource
	sourcetypes.KnSourceClient
}

type GitHubSourceParamsFactory interface {
	GitHubSource
	sourcetypes.ParamsFactory

	CreateGitHubSourceParams() *GitHubSourceParams
}

type GitHubSourceClientFactory interface {
	GitHubSource
	sourcetypes.ClientFactory

	CreateGitHubSourceClient() GitHubSourceClient
}

type GitHubSourceCommandFactory interface {
	GitHubSource
	sourcetypes.CommandFactory
}

type GitHubSourceFlagsFactory interface {
	GitHubSource
	sourcetypes.FlagsFactory
}

type GitHubSourceRunEFactory interface {
	GitHubSource
	sourcetypes.RunEFactory

	GitHubSourceClient() GitHubSourceClient
}
