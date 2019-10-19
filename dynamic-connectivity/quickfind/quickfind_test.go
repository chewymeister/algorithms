package quickfind_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chewymeister/algorithms/dynamic-connectivity/quickfind"
)

var _ = Describe("Quickfind", func() {
	var (
		finder *quickfind.QuickFinder
	)

	BeforeEach(func() {
		finder = quickfind.New(10)
	})

	Context("union", func() {
		It("should connect a single pair of points", func() {
			Expect(finder.Union(0, 1)).To(Succeed())
			Expect(finder.Connected(0, 1)).To(BeTrue())
			Expect(finder.Connected(2, 3)).To(BeFalse())
		})

		It("should connect multiple pairs of points", func() {
			Expect(finder.Union(0, 1)).To(Succeed())
			Expect(finder.Union(2, 3)).To(Succeed())
			Expect(finder.Connected(0, 1)).To(BeTrue())
			Expect(finder.Connected(2, 3)).To(BeTrue())
		})

		It("should connect multiple common pairs of points", func() {
			Expect(finder.Union(0, 1)).To(Succeed())
			Expect(finder.Union(2, 3)).To(Succeed())
			Expect(finder.Union(1, 2)).To(Succeed())
			Expect(finder.Connected(0, 1)).To(BeTrue())
			Expect(finder.Connected(2, 3)).To(BeTrue())
			Expect(finder.Connected(0, 3)).To(BeTrue())
		})

		It("should error if Union parameter is greater than length of finder", func() {
			finder := quickfind.New(2)
			Expect(finder.Union(0, 2)).To(MatchError("could not make union, index is greater than length of initialized finder"))
		})
	})
})
