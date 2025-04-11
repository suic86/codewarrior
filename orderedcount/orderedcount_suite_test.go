package orderedcount_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOrderedcount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Orderedcount Suite")
}
