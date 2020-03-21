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
	"github.com/maximilien/kn-source-github/pkg/client"
	"github.com/maximilien/kn-source-github/pkg/types"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type gitHubClientFactory struct {
	gitHubSourceClient  types.GitHubSourceClient
	gitHubParamsFactory types.GitHubSourceParamsFactory
}

func NewGitHubSourceClientFactory(gitHubParamsFactory types.GitHubSourceParamsFactory) types.GitHubSourceClientFactory {
	return &gitHubClientFactory{
		gitHubParamsFactory: gitHubParamsFactory,
	}
}

func (f *gitHubClientFactory) CreateKnSourceClient() sourcetypes.KnSourceClient {
	return f.CreateGitHubSourceClient()
}

func (f *gitHubClientFactory) CreateGitHubSourceClient() types.GitHubSourceClient {
	if f.gitHubSourceClient == nil {
		f.initGitHubSourceClient()
	}
	return f.gitHubSourceClient
}

func (f *gitHubClientFactory) GitHubSourceParams() *types.GitHubSourceParams {
	return f.gitHubParamsFactory.GitHubSourceParams()
}

func (f *gitHubClientFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	return f.gitHubParamsFactory.KnSourceParams()
}

// Private

func (f *gitHubClientFactory) initGitHubSourceClient() {
	if f.gitHubSourceClient == nil {
		f.gitHubSourceClient = client.NewGitHubSourceClient(f.gitHubParamsFactory)
	}
}
