package cyberq_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCyberq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cyberq Suite")
}
