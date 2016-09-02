package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQrGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "QrGenerator Suite")
}
