package mostfrequentdays_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMostfrequentdays(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mostfrequentdays Suite")
}
