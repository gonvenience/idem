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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gonvenience/idem"
)

var _ = Describe("Rename", func() {
	Context("Detect Rename", func() {
		It("should", func() {
			var changes = NewDocumentChanges(
				[]*RenameCandidate{
					NewRenameCandidate("bar", path("/foo/bar"), node()),
				},
				[]*RenameCandidate{
					NewRenameCandidate("bAr", path("/foo/bAr"), node()),
				},
			)

			Expect(DetectRenames(changes, nil)).To(Succeed())

			Expect(changes.ModifiedPairs()).To(HaveLen(1))
			Expect(changes.ModifiedPairs()[0].From.Name()).To(Equal("bar"))
			Expect(changes.ModifiedPairs()[0].To.Name()).To(Equal("bAr"))
		})
	})
})
