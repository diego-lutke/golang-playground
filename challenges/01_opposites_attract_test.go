// adapted from: https://www.codewars.com/kata/555086d53eac039a2a000083/train/go

package challenges

import (
	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var _ = g.Describe("Tests", func() {
	g.It("Basic Tests", func() {
		o.Expect(LoveFunc(1, 4)).To(o.Equal(true))
		o.Expect(LoveFunc(2, 2)).To(o.Equal(false))
		o.Expect(LoveFunc(0, 1)).To(o.Equal(true))
		o.Expect(LoveFunc(0, 0)).To(o.Equal(false))
	})
})
