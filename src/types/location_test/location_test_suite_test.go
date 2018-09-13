package location_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLocationTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LocationTest Suite")
}
