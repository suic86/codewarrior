package uniq_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUniq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uniq Suite")
}
