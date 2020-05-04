package mixbonacci_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMixbonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mixbonacci Suite")
}
