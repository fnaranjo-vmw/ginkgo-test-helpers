package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Using helper functions in tests", func() {
	Context("Ginkgo", func() {
		When("assertions are used in the helper function", func() {
			It("prints USELESS line number in colorized text and summaries", func() {
				ginkgoHelper1("some-value")
			})
		})
		When("assertions are used in the test body", func() {
			It("prints USEFUL line number in colorized text and summaries", func() {
				err := ginkgoHelper2("some-value")
				if err != nil {
					Fail("Something is failing")
				}
			})
		})
	})
	Context("Gomega", func() {
		When("assertions are used in the helper function", func() {
			It("prints USELESS line number but USEFUL automatic text in colorized text and summaries", func() {
				gomegaHelper1("some-value")
			})
		})
		When("assertions are used in the test body", func() {
			It("prints USEFUL line number but MISSES automatic text in colorized text and summaries", func() {
				err := gomegaHelper2("some-value")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		When("the helper function returns an GomegaMatcher", func() {
			It("prints USEFUL line number and USEFUL automatic text in colorized text and summaries. Getting the best of both worlds", func() {
				val := "some-value"
				Expect(val).To(gomegaHelper3(val))
				// Using better functions names makes the code easy to read
				// Expect(val).To(beASupportedValue(val))
			})
		})
	})
})
