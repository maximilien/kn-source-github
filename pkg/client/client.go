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

	knerrors "knative.dev/client/pkg/errors"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "knative.dev/eventing-contrib/github/pkg/apis/sources/v1alpha1"

	clientv1alpha1 "knative.dev/eventing-contrib/github/pkg/client/clientset/versioned/typed/sources/v1alpha1"
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

// Namespace for this client
func (client *ghSourceClient) Namespace() string {
	return client.knSourceClient.Namespace()
}

// SourcesClient the client to access the Knative sources
func (client *ghSourceClient) SourcesClient() clientv1alpha1.SourcesV1alpha1Interface {
	return client.knSourceClient.SourcesClient()
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

// GetGHSource is used to create and initialize an GHSource
func (client *ghSourceClient) GetGHSource(name string) (*v1alpha1.GitHubSource, error) {
	ghSource, err := client.SourcesClient().GitHubSources(client.namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, knerrors.GetError(err)
	}
	return ghSource, nil
}

// CreateGHSource is used to create and initialize an GHSource
func (client *ghSourceClient) CreateGHSource(ghSource *v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error) {
	ghSource, err := client.SourcesClient().GitHubSources(client.namespace).Create(ghSource)
	if err != nil {
		return nil, knerrors.GetError(err)
	}
	return ghSource, nil
}

// UpdateGHSource is used to update a GHSource
func (client *ghSourceClient) UpdateGHSource(ghSource *v1alpha1.GitHubSource) (*v1alpha1.GitHubSource, error) {
	ghSource, err := client.SourcesClient().GitHubSources(client.namespace).Update(ghSource)
	if err != nil {
		return nil, knerrors.GetError(err)
	}
	return ghSource, nil
}

// DeleteGHSource is used to delete an GHSource
func (client *ghSourceClient) DeleteGHSource(name string) error {
	err := client.SourcesClient().GitHubSources(client.namespace).Delete(name, &v1.DeleteOptions{})
	if err != nil {
		return knerrors.GetError(err)
	}

	return nil
}
