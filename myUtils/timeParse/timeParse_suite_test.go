package timeParse_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTimeParse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TimeParse Suite")
}
