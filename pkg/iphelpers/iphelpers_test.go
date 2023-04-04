package iphelpers

import (
	"net"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIPHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmd")
}

var _ = Describe("CompareIPs operations", func() {
	It("compares IPv4.To4() addresses", func() {
		left := net.ParseIP("192.168.0.0")
		right := net.ParseIP("192.169.0.0")
		Expect(CompareIPs(left.To4(), right.To4())).To(Equal(-1))
	})

	It("compares IPv4.To16() addresses", func() {
		left := net.ParseIP("192.169.0.0")
		right := net.ParseIP("192.168.0.0")
		Expect(CompareIPs(left.To16(), right.To16())).To(Equal(1))
	})

	It("compares IPv4 mixed addresses", func() {
		left := net.ParseIP("192.168.0.0")
		right := net.ParseIP("192.168.0.0")
		Expect(CompareIPs(left.To16(), right.To4())).To(Equal(0))
	})

	It("compares IPv6 addresses when left is smaller than right", func() {
		left := net.ParseIP("2000::")
		right := net.ParseIP("2000::1")
		Expect(CompareIPs(left, right)).To(Equal(-1))
	})

	It("compares IPv6 addresses when left is larger than right", func() {
		left := net.ParseIP("2000::1")
		right := net.ParseIP("2000::")
		Expect(CompareIPs(left, right)).To(Equal(1))
	})

	It("compares IPv6 addresses when left == right", func() {
		left := net.ParseIP("2000::1")
		right := net.ParseIP("2000::1")
		Expect(CompareIPs(left, right)).To(Equal(0))
	})
})
