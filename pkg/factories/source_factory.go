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

	sourcefactories "github.com/maximilien/kn-source-pkg/pkg/factories"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type gitHubSourceFactory struct {
	gitHubSourceParams *types.GitHubSourceParams
	gitHubSourceClient types.GitHubSourceClient

	knSourceFactory sourcetypes.KnSourceFactory
}

func NewGitHubSourceFactory() types.GitHubSourceFactory {
	return &gitHubSourceFactory{
		gitHubSourceParams: nil,
		gitHubSourceClient: nil,
		knSourceFactory:    sourcefactories.NewDefaultKnSourceFactory(),
	}
}

func (f *gitHubSourceFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams.KnSourceParams
}

func (f *gitHubSourceFactory) GitHubSourceParams() *types.GitHubSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams
}

func (f *gitHubSourceFactory) GitHubSourceClient() types.GitHubSourceClient {
	return f.gitHubSourceClient
}

func (f *gitHubSourceFactory) CreateKnSourceParams() *sourcetypes.KnSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams.KnSourceParams
}

func (f *gitHubSourceFactory) CreateGitHubSourceParams() *types.GitHubSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams
}

func (f *gitHubSourceFactory) CreateKnSourceClient(namespace string) sourcetypes.KnSourceClient {
	return f.CreateGitHubSourceClient(namespace)
}

func (f *gitHubSourceFactory) CreateGitHubSourceClient(namespace string) types.GitHubSourceClient {
	if f.gitHubSourceClient == nil {
		f.initGitHubSourceClient(namespace)
	}
	return f.gitHubSourceClient
}

// Private

func (f *gitHubSourceFactory) initGitHubSourceClient(namespace string) {
	if f.gitHubSourceClient == nil {
		f.gitHubSourceClient = client.NewGitHubSourceClient(f, namespace)
	}
}

// Private

func (f *gitHubSourceFactory) initGitHubSourceParams() {
	f.gitHubSourceParams = &types.GitHubSourceParams{
		KnSourceParams: f.knSourceFactory.CreateKnSourceParams(),
	}
	f.gitHubSourceParams.KnSourceParams.Initialize()
}
