// Copyright 2025 The Homeport Team
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

package idem_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gonvenience/ytbx"
	yamlv3 "go.yaml.in/yaml/v3"
)

func TestIdem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Idem Suite")
}

func path(in string) *ytbx.Path {
	GinkgoHelper()

	path, err := ytbx.ParsePathStringUnsafe(in)
	Expect(err).ToNot(HaveOccurred())
	return &path
}

func node() *yamlv3.Node {
	GinkgoHelper()
	return &yamlv3.Node{}
}
