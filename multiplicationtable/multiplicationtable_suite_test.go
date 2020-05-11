package multiplicationtable_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultiplicationtable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multiplicationtable Suite")
}
