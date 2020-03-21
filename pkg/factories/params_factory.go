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
	"github.com/maximilien/kn-source-github/pkg/types"

	sourcefactories "github.com/maximilien/kn-source-pkg/pkg/factories"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type gitHubSourceParamsFactory struct {
	defaultParamsFactory sourcetypes.ParamsFactory
	gitHubSourceParams   *types.GitHubSourceParams
}

func NewGitHubSourceParamsFactory() types.GitHubSourceParamsFactory {
	return &gitHubSourceParamsFactory{
		defaultParamsFactory: sourcefactories.NewDefaultParamsFactory(),
	}
}

func (f *gitHubSourceParamsFactory) KnSourceParams() *sourcetypes.KnSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams.KnSourceParams
}

func (f *gitHubSourceParamsFactory) GitHubSourceParams() *types.GitHubSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams
}

func (f *gitHubSourceParamsFactory) CreateKnSourceParams() *sourcetypes.KnSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams.KnSourceParams
}

func (f *gitHubSourceParamsFactory) CreateGitHubSourceParams() *types.GitHubSourceParams {
	if f.gitHubSourceParams == nil {
		f.initGitHubSourceParams()
	}
	return f.gitHubSourceParams
}

// Private

func (f *gitHubSourceParamsFactory) initGitHubSourceParams() {
	f.gitHubSourceParams = &types.GitHubSourceParams{
		KnSourceParams: f.defaultParamsFactory.CreateKnSourceParams(),
	}
	f.gitHubSourceParams.KnSourceParams.Initialize()
}
