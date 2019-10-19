package quickunion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chewymeister/algorithms/dynamic-connectivity/quickunion"
)

var _ = Describe("Quickfind", func() {
	var (
		unioner *quickunion.QuickUnioner
	)

	BeforeEach(func() {
		unioner = quickunion.New(10)
	})

	Context("union", func() {
		It("should connect a single pair of points", func() {
			Expect(unioner.Union(0, 1)).To(Succeed())
			Expect(unioner.Connected(0, 1)).To(BeTrue())
			Expect(unioner.Connected(2, 3)).To(BeFalse())
		})

		It("should connect multiple pairs of points", func() {
			Expect(unioner.Union(0, 1)).To(Succeed())
			Expect(unioner.Union(2, 3)).To(Succeed())
			Expect(unioner.Connected(0, 1)).To(BeTrue())
			Expect(unioner.Connected(2, 3)).To(BeTrue())
		})

		It("should connect multiple common pairs of points", func() {
			Expect(unioner.Union(0, 1)).To(Succeed())
			Expect(unioner.Union(2, 3)).To(Succeed())
			Expect(unioner.Union(1, 2)).To(Succeed())
			Expect(unioner.Connected(0, 1)).To(BeTrue())
			Expect(unioner.Connected(2, 3)).To(BeTrue())
			Expect(unioner.Connected(0, 3)).To(BeTrue())
		})

		It("should error if Union parameter is greater than length of unioner", func() {
			unioner := quickunion.New(2)
			Expect(unioner.Union(0, 2)).To(MatchError("could not make union, point is greater than length of initialized unioner"))
		})
	})
})
