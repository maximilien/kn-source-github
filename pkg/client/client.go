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

type ghSourceClient struct {
	namespace      string
	ghSourceParams *types.GHSourceParams
	knSourceClient sourcetypes.KnSourceClient
}

func NewGHSourceClient(ghSourceParams *types.GHSourceParams, namespace string) types.GHSourceClient {
	return &ghSourceClient{
		namespace:      namespace,
		ghSourceParams: ghSourceParams,
		knSourceClient: sourceclient.NewKnSourceClient(ghSourceParams.KnSourceParams, namespace),
	}
}

func (client *ghSourceClient) Namespace() string {
	return client.knSourceClient.Namespace()
}

func (client *ghSourceClient) KnSourceClient() sourcetypes.KnSourceClient {
	return client
}

func (client *ghSourceClient) KnSourceParams() *sourcetypes.KnSourceParams {
	return client.GHSourceParams().KnSourceParams
}

func (client *ghSourceClient) GHSourceParams() *types.GHSourceParams {
	return client.ghSourceParams
}
