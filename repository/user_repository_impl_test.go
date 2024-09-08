package repository

import (
	"gopher-bank/model"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	userRepository = NewUserRepository()
)

func TestGopherBank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GopherBank Suite")
}

var _ = Describe("./Repository/UserRepositoryImpl", func() {
	It("FindAll return user when user was create", func() {
		user := &model.User{
			Username: "test",
			Email: "test@email.com",
			Password: "test123",
		}
		userRepository.Save(user)
		users, _ := userRepository.FindAll()

		Expect(&users[0]).To(Equal(user))
	})
})
