package macdive_test

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MacDiveTest struct {
}

var _ = Suite(&MacDiveTest{})
