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
	v1alpha1 "knative.dev/eventing-contrib/github/pkg/apis/sources/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// GitHubSourceBuilder is for building the source
type GitHubSourceBuilder struct {
	ghSource *v1alpha1.GitHubSource
}

// NewGitHubSourceBuilder for building ApiServer source object
func NewGitHubSourceBuilder(name string) *GitHubSourceBuilder {
	return &GitHubSourceBuilder{
		ghSource: &v1alpha1.GitHubSource{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	}
}

// // BootstrapServers to set the value of BootstrapServers
// func (b *GitHubSourceBuilder) BootstrapServers(server string) *GitHubSourceBuilder {
// 	b.ghSource.Spec.BootstrapServers = server
// 	return b
// }

// // Topics to set the value of Topics
// func (b *GitHubSourceBuilder) Topics(topics string) *GitHubSourceBuilder {
// 	b.ghSource.Spec.Topics = topics
// 	return b
// }

// // ConsumerGroup to set the value of ConsumerGroup
// func (b *GitHubSourceBuilder) ConsumerGroup(consumerGroup string) *GitHubSourceBuilder {
// 	b.ghSource.Spec.ConsumerGroup = consumerGroup
// 	return b
// }

// Sink or destination of the source
func (builder *GitHubSourceBuilder) Sink(sink *duckv1.Destination) *GitHubSourceBuilder {
	builder.ghSource.Spec.Sink = sink
	return builder
}

// Build the GitHubSource object
func (b *GitHubSourceBuilder) Build() *v1alpha1.GitHubSource {
	return b.ghSource
}
