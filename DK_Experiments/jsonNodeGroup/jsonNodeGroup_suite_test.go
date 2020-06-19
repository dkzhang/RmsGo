package jsonNodeGroup_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJsonNodeGroup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JsonNodeGroup Suite")
}
