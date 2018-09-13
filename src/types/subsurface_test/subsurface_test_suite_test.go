package subsurface_test

import (
	"github.com/onsi/ginkgo/reporters"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSubsurfaceTest(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "SubsurfaceTest Suite")
	junitReporter := reporters.NewJUnitReporter("junit.xml")
    RunSpecsWithDefaultAndCustomReporters(t, "SubsurfaceTest Suite", []Reporter{junitReporter})
}
