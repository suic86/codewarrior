package uniq_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUniq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uniq Suite")
}
