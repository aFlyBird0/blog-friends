package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generate", func() {
	When("input is valid", func() {
		group, shortcodeGenerated, err := Generate(inputForTestWithCustomGroup, TplFile)
		if err != nil {
			return
		}
		It("should return valid result", func() {
			Expect(err).To(BeNil())
			Expect(group).To(Equal(customGroupForTest))
			Expect(shortcodeGenerated).To(Equal(shortCodeForTest))
		})
	})
})
