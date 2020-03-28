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
	sourceclient "github.com/maximilien/kn-source-pkg/pkg/client"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"
)

type gitHubSourceClient struct {
	namespace          string
	gitHubSourceParams *types.GitHubSourceParams
	knSourceClient     sourcetypes.KnSourceClient
}

func NewGitHubSourceClient(gitHubSourceParams *types.GitHubSourceParams, namespace string) types.GitHubSourceClient {
	return &gitHubSourceClient{
		namespace:          namespace,
		gitHubSourceParams: gitHubSourceParams,
		knSourceClient:     sourceclient.NewKnSourceClient(gitHubSourceParams.KnSourceParams, namespace),
	}
}

func (client *gitHubSourceClient) Namespace() string {
	return client.knSourceClient.Namespace()
}

func (client *gitHubSourceClient) KnSourceClient() sourcetypes.KnSourceClient {
	return client
}

func (client *gitHubSourceClient) KnSourceParams() *sourcetypes.KnSourceParams {
	return client.GitHubSourceParams().KnSourceParams
}

func (client *gitHubSourceClient) GitHubSourceParams() *types.GitHubSourceParams {
	return client.gitHubSourceParams
}
