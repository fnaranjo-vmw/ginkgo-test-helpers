package main

import (
	"fmt"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func ginkgoHelper1(val string) {
	ginkgo.Fail("Something is failing")
}

func ginkgoHelper2(val string) error {
	return fmt.Errorf("Something is failing")
}

func gomegaHelper1(val string) {
	Expect(val).To(Equal("never ever fail"))
}

func gomegaHelper2(val string) error {
	if val != "never ever fail" {
		return fmt.Errorf("expecting %q but got %q", "never ever fail", val)
	}
	return nil
}

func gomegaHelper3() types.GomegaMatcher {
	return Equal("never ever fail")
}
