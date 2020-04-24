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
	"errors"
	"fmt"

	"github.com/maximilien/kn-source-github/pkg/client"
	"github.com/maximilien/kn-source-github/pkg/types"

	sourcetypes "github.com/maximilien/kn-source-pkg/pkg/types"

	"github.com/spf13/cobra"
)

type ghRunEFactory struct {
	ghSourceFactory types.GHSourceFactory
}

func NewGHRunEFactory(ghSourceFactory types.GHSourceFactory) types.GHRunEFactory {
	return &ghRunEFactory{
		ghSourceFactory: ghSourceFactory,
	}
}

func (f *ghRunEFactory) CreateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		ghSourceClient := f.GHSourceClient(namespace)

		if len(args) != 1 {
			return errors.New("requires the NAME of the source to create as single argument")
		}

		name := args[0]

		dynamicClient, err := f.KnSourceParams().KnParams.NewDynamicClient(namespace)
		if err != nil {
			return err
		}

		objectRef, err := f.KnSourceParams().SinkFlag.ResolveSink(dynamicClient, namespace)
		if err != nil {
			return fmt.Errorf("cannot create GitHub source '%s' in namespace '%s' because: %s",
				name, namespace, err.Error())
		}

		builder := client.NewGitHubSourceBuilder(name).
			OrgRepo(ghSourceClient.GHSourceParams().Org, ghSourceClient.GHSourceParams().Repo).
			APIURL(ghSourceClient.GHSourceParams().APIURL).
			AccessToken(ghSourceClient.GHSourceParams().AccessToken).
			SecretToken(ghSourceClient.GHSourceParams().SecretToken).
			Sink(objectRef)

		_, err = ghSourceClient.CreateGHSource(builder.Build())
		if err != nil {
			return fmt.Errorf(
				"cannot create GitHub source '%s' in namespace '%s' because: %s",
				name, namespace, err.Error())
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "GitHub source '%s' created in namespace '%s'.\n", name, namespace)
		}

		return nil
	}
}

func (f *ghRunEFactory) DeleteRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		ghSourceClient := f.GHSourceClient(namespace)

		if len(args) != 1 {
			return errors.New("requires the NAME of the source to delete as single argument")
		}

		name := args[0]

		err = ghSourceClient.DeleteGHSource(name)
		if err != nil {
			return fmt.Errorf(
				"cannot delete GitHub source '%s' in namespace '%s' because: %s",
				name, namespace, err.Error())
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "GitHub source '%s' deleted in namespace '%s'.\n", name, namespace)
		}

		return nil
	}
}

func (f *ghRunEFactory) UpdateRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		ghSourceClient := f.GHSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v, sink: %s\n", cmd.Name(), args, ghSourceClient, ghSourceClient.KnSourceParams().SinkFlag)
		return nil
	}
}

func (f *ghRunEFactory) DescribeRunE() sourcetypes.RunE {
	return func(cmd *cobra.Command, args []string) error {
		namespace, err := f.KnSourceParams().GetNamespace(cmd)
		if err != nil {
			return err
		}

		ghSourceClient := f.GHSourceClient(namespace)
		fmt.Printf("%s RunE function called for GitHub source: args: %#v, client: %#v, sink: %s\n", cmd.Name(), args, ghSourceClient, ghSourceClient.KnSourceParams().SinkFlag)
		return nil
	}
}
