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

	v1alpha1 "knative.dev/eventing-contrib/github/pkg/apis/sources/v1alpha1"
)

type GHSourceClient interface {
	sourcetypes.KnSourceClient
	GHSourceParams() *GHSourceParams

	InitGHSource(ghSource *v1alpha1.GitHubSource) error
}

type GHSourceFactory interface {
	sourcetypes.KnSourceFactory

	GHSourceParams() *GHSourceParams
	GHSourceClient() GHSourceClient

	CreateGHSourceParams() *GHSourceParams
	CreateGHSourceClient(namespace string) GHSourceClient
}

type GHCommandFactory interface {
	sourcetypes.CommandFactory

	GHSourceFactory() GHSourceFactory
}

type GHFlagsFactory interface {
	sourcetypes.FlagsFactory

	GHSourceFactory() GHSourceFactory
}

type GHRunEFactory interface {
	sourcetypes.RunEFactory

	GHSourceFactory() GHSourceFactory
	GHSourceClient(namespace string) GHSourceClient
}
