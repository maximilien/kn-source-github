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

package factories

import (
	"testing"

	"gotest.tools/assert"

	"github.com/maximilien/kn-source-github/pkg/types"
)

func TestNewGHRunEFactory(t *testing.T) {
	runEFactory := createGHRunEFactory()

	assert.Assert(t, runEFactory != nil)
}

func TestCreateRunE(t *testing.T) {
	runEFactory := createGHRunEFactory()

	createRunE := runEFactory.CreateRunE()
	assert.Assert(t, createRunE != nil)
}

func TestDeleteRunE(t *testing.T) {
	runEFactory := createGHRunEFactory()

	deleteRunE := runEFactory.DeleteRunE()
	assert.Assert(t, deleteRunE != nil)
}

func TestUpdateRunE(t *testing.T) {
	runEFactory := createGHRunEFactory()

	updateRunE := runEFactory.UpdateRunE()
	assert.Assert(t, updateRunE != nil)
}

func TestDescribeRunE(t *testing.T) {
	runEFactory := createGHRunEFactory()

	describeRunE := runEFactory.DescribeRunE()
	assert.Assert(t, describeRunE != nil)
}

// Private

func createGHRunEFactory() types.GHRunEFactory {
	ghSourceFactory := NewGHSourceFactory()
	return NewGHRunEFactory(ghSourceFactory)
}
