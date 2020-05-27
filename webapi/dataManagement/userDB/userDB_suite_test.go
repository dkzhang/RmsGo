package userDB_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUserDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserDB Suite")
}
