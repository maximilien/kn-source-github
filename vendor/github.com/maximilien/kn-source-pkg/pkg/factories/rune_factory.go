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
	"fmt"

	"github.com/maximilien/kn-source-pkg/pkg/commands"

	"github.com/spf13/cobra"
)

type DefautRunEFactory struct{}

func NewDefaultRunEFactory() commands.RunEFactory {
	return &DefautRunEFactory{}
}

func (f *DefautRunEFactory) CreateRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("create RunE called: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *DefautRunEFactory) DeleteRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("delete RunE called: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *DefautRunEFactory) UpdateRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("update RunE called: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}

func (f *DefautRunEFactory) DescribeRunE() commands.RunE {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("describe RunE called: %s, args: %#v\n", cmd.Name(), args)
		return nil
	}
}
