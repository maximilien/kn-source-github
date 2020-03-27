// Copyright © 2019 The Knative Authors
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

package client

import (
	"github.com/maximilien/kn-source-github/pkg/types"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type gitHubSourceClient struct {
	knSourceClient sourcetypes.KnSourceClient

	gitHubSourceFactory types.GitHubSourceFactory
}

func NewGitHubSourceClient(gitHubSourceFactory types.GitHubSourceFactory, namespace string) types.GitHubSourceClient {
	return &gitHubSourceClient{
		gitHubSourceFactory: gitHubSourceFactory,
		knSourceClient:      gitHubSourceFactory.CreateKnSourceClient(namespace),
	}
}

func (client *gitHubSourceClient) Namespace() string {
	return client.knSourceClient.Namespace()
}

func (client *gitHubSourceClient) KnSourceClient() sourcetypes.KnSourceClient {
	return client
}

func (client *gitHubSourceClient) KnSourceParams() *sourcetypes.KnSourceParams {
	return client.gitHubSourceFactory.GitHubSourceParams().KnSourceParams
}

func (client *gitHubSourceClient) GitHubSourceParams() *types.GitHubSourceParams {
	return client.gitHubSourceFactory.GitHubSourceParams()
}
