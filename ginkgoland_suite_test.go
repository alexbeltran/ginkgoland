package ginkgoland_test

import (
	"github.com/alexbeltran/ginkgoland"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoland(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "Ginkgoland Suite")
	//RunSpecsWithDefaultAndCustomReporters(t, "Ginkgoland Suite", []Reporter{&ginkgoland.Reporter{}})
	RunSpecsWithCustomReporters(t, "MyGinkgoland Suite Suite Life", []Reporter{&ginkgoland.Reporter{T: t}})
}
