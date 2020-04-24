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
	"testing"

	"gotest.tools/assert"

	"github.com/maximilien/kn-source-github/pkg/types"
)

// TODO: fix this
func _TestNewGHSourceClient(t *testing.T) {
	ghSourceClient := NewGHSourceClient(&types.GHSourceParams{}, "fake-namespace")
	assert.Assert(t, ghSourceClient != nil)
}

// TODO: fix this
func _TestGHSourceParams(t *testing.T) {
	fakeGHSourceParams := &types.GHSourceParams{}
	ghSourceClient := NewGHSourceClient(fakeGHSourceParams, "fake-namespace")
	assert.Equal(t, ghSourceClient.GHSourceParams(), fakeGHSourceParams)
}

// TODO: fix this
func _TestNamespace(t *testing.T) {
	ghSourceClient := NewGHSourceClient(&types.GHSourceParams{}, "fake-namespace")
	assert.Equal(t, ghSourceClient.Namespace(), "fake-namespace")
}
