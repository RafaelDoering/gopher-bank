package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGopherBank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GopherBank Suite")
}
