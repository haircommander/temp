package utils_test

import (
	"github.com/kubernetes-sigs/cri-o/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// The actual test suite
var _ = t.Describe("Filesystem", func() {
	t.Describe("GetDiskUsageStats", func() {
		It("should succeed at the current working directory", func() {
			// Given
			// When
			bytes, inodes, err := utils.GetDiskUsageStats(".")

			// Then
			Expect(err).To(BeNil())
			Expect(bytes).To(SatisfyAll(BeNumerically(">", 0)))
			Expect(inodes).To(SatisfyAll(BeNumerically(">", 0)))
		})

		It("should fail on invalid path", func() {
			// Given
			// When
			bytes, inodes, err := utils.GetDiskUsageStats("/not-existing")

			// Then
			Expect(err).NotTo(BeNil())
			Expect(bytes).To(BeEquivalentTo(0))
			Expect(inodes).To(BeEquivalentTo(0))
		})
	})
})
