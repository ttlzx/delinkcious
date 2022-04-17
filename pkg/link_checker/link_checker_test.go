package link_checker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Link checker tests", func() {
	It("should not return error for a valid url", func() {
		err := CheckLink("https://github.com/ttlzx/delinkcious")
		Ω(err).Should(BeNil())
	})

	It("should not return error for non-existent url", func() {
		err := CheckLink("https://github.com/ttlzx/delinkcious/no-such-url")
		Ω(err).ShouldNot(BeNil())
	})
})
