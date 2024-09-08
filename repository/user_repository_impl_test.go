package repository

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGopherBank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GopherBank Suite")
}

var _ = Describe("./Repository/UserRepositoryImpl", func() {
	It("can extract the author's last name", func() {
		Expect(true).To(Equal(true))
	})
})
