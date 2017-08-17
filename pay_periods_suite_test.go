package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPayPeriods(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PayPeriods Suite")
}
