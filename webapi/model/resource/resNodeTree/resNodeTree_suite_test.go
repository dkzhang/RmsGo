package resNodeTree_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResNodeTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResNodeTree Suite")
}
