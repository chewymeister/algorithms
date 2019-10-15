package quickfind_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuickfind(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Quickfind Suite")
}
