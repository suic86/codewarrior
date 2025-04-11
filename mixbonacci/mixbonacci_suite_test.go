package mixbonacci_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMixbonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mixbonacci Suite")
}
