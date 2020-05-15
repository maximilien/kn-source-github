// Copyright © 2020 The Knative Authors
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

	"github.com/maximilien/kn-source-pkg/pkg/types"

	clientv1alpha1 "knative.dev/eventing-contrib/github/pkg/client/clientset/versioned/typed/sources/v1alpha1"
)

type knSourceClient struct {
	knSourceParams *types.KnSourceParams
	namespace      string
	sourcesClient  clientv1alpha1.SourcesV1alpha1Interface
}

// NewKnSourceClient creates a new KnSourceClient with parameters and namespace
func NewKnSourceClient(knSourceParams *types.KnSourceParams, namespace string) types.KnSourceClient {
	sourcesClient, err := newSourcesClient(knSourceParams)
	if err != nil {
		panic(fmt.Sprintf("Could not create GitHub sources client, error: %s", err.Error()))
	}

	return &knSourceClient{
		knSourceParams: knSourceParams,
		namespace:      namespace,
		sourcesClient:  sourcesClient,
	}
}

// KnSourceParams returns the client's KnSourceParams
func (client *knSourceClient) KnSourceParams() *types.KnSourceParams {
	return client.knSourceParams
}

// Namespace returns the client's namespace
func (client *knSourceClient) Namespace() string {
	return client.namespace
}

// SourcesClient the client to access the Knative sources
func (client *knSourceClient) SourcesClient() clientv1alpha1.SourcesV1alpha1Interface {
	return client.sourcesClient
}

// Private

func newSourcesClient(knSourceParams *types.KnSourceParams) (*clientv1alpha1.SourcesV1alpha1Client, error) {
	restConfig, err := knSourceParams.RestConfig()
	if err != nil {
		return nil, err
	}

	return clientv1alpha1.NewForConfig(restConfig)
}
