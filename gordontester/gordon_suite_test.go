package gordontester_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGordon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gordon Suite")
}
