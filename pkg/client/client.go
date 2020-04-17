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
	"fmt"

	"github.com/maximilien/kn-source-github/pkg/types"

	sourceclient "github.com/maximilien/kn-source-pkg/pkg/client"
	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	knerrors "knative.dev/client/pkg/errors"

	v1alpha1 "knative.dev/eventing-contrib/github/pkg/apis/sources/v1alpha1"
	clientv1alpha1 "knative.dev/eventing-contrib/github/pkg/client/clientset/versioned/typed/sources/v1alpha1"
)

type ghSourceClient struct {
	namespace      string
	ghSourceParams *types.GHSourceParams
	knSourceClient sourcetypes.KnSourceClient
	sourcesClient  clientv1alpha1.SourcesV1alpha1Interface
}

func NewGHSourceClient(ghSourceParams *types.GHSourceParams, namespace string) types.GHSourceClient {
	sourcesClient, err := newSourcesClient(ghSourceParams)
	if err != nil {
		panic(fmt.Sprintf("Could not create GitHub sources client, error: %s", err.Error()))
	}

	return &ghSourceClient{
		namespace:      namespace,
		ghSourceParams: ghSourceParams,
		knSourceClient: sourceclient.NewKnSourceClient(ghSourceParams.KnSourceParams, namespace),
		sourcesClient:  sourcesClient,
	}
}

// Namespace for this client
func (client *ghSourceClient) Namespace() string {
	return client.knSourceClient.Namespace()
}

// KnSourceClient interface for this client
func (client *ghSourceClient) KnSourceClient() sourcetypes.KnSourceClient {
	return client
}

// KnSourceParams for common Kn source parameters
func (client *ghSourceClient) KnSourceParams() *sourcetypes.KnSourceParams {
	return client.GHSourceParams().KnSourceParams
}

// GHSourceParams for GitHub source specific parameters
func (client *ghSourceClient) GHSourceParams() *types.GHSourceParams {
	return client.ghSourceParams
}

// InitGHSource is used to create and initialize an instance of ApiServerSource
func (client *ghSourceClient) InitGHSource(ghSource *v1alpha1.GitHubSource) error {
	_, err := client.sourcesClient.GitHubSources(client.namespace).Create(ghSource)
	if err != nil {
		return knerrors.GetError(err)
	}

	return nil
}

// Private

func newSourcesClient(ghSourceParams *types.GHSourceParams) (*clientv1alpha1.SourcesV1alpha1Client, error) {
	restConfig, err := ghSourceParams.KnSourceParams.RestConfig()
	if err != nil {
		return nil, err
	}

	return clientv1alpha1.NewForConfig(restConfig)
}
