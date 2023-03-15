package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generate", func() {
	When("input is valid", func() {
		It("should return valid result", func() {
			group, shortcodeGenerated, err := Generate(inputForTestWithCustomGroup, TplFile)

			Expect(err).To(BeNil())
			Expect(group).To(Equal(customGroupForTest))
			Expect(shortcodeGenerated).To(Equal(shortCodeForTest))
		})
	})
})
