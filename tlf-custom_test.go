package goblin

import (
	"testing"
)

func Testtlf(tes *testing.T) {
	g := Goblin(tes)

	g.Describe("Catching", func() {
		g.It("No panic", func() {
		})

		g.It("Expected panic", func() {
			g.ExpectPanic()
			panic("Expect")
		})

		g.It("Fail on no expected panic", func() {
			g.ExpectFail()
			g.ExpectPanic()
		})

		g.It("Fail on unexpected panic", func() {
			g.ExpectFail()
			panic("Suddenly, panicing!")
		})
	})

	g.Describe("Custom goblin assert", func() {
		g.Describe("Working", func() {
			g.Describe("Normal", func() {
				g.It("Eql", func() {
					g.Assert("Hello").Eql("Hello")
				})

				g.It("IsTrue", func() {
					g.Assert(true).IsTrue()
				})

				g.It("IsFalse", func() {
					g.Assert(false).IsFalse()
				})

				g.It("IsOK", func() {
					g.Assert(g).IsOK()
				})

				g.It("IsType", func() {
					g.Assert("").HasSameType("")
				})
			})

			g.Describe("Not", func() {
				g.It("Eql", func() {
					g.Assert("Hello").Not.Eql("Lol")
					g.Assert("Hello").Not.Eql("hello")
				})

				g.It("IsTrue", func() {
					g.Assert(false).Not.IsTrue()
					g.Assert("").Not.IsTrue()
				})

				g.It("IsFalse", func() {
					g.Assert(true).Not.IsFalse()
					g.Assert("a").Not.IsFalse()
				})

				g.It("IsOK", func() {
					g.Assert(nil).Not.IsOK()
				})

				g.It("IsType", func() {
					g.Assert(1).Not.HasSameType("")
				})
			})
		})

		g.Describe("Failing", func() {
			g.Describe("Normal", func() {
				g.It("Eql", func() {
					g.ExpectFail()
					g.Assert("Hello").Eql("hello")
				})

				g.It("IsTrue", func() {
					g.ExpectFail()
					g.Assert(false).IsTrue()
				})

				g.It("IsFalse", func() {
					g.ExpectFail()
					g.Assert(true).IsFalse()
				})

				g.It("IsOK", func() {
					g.ExpectFail()
					g.Assert(nil).IsOK()
				})

				g.It("IsType", func() {
					g.ExpectFail()
					g.Assert(1).HasSameType("")
				})
			})

			g.Describe("Not", func() {
				g.It("Eql", func() {
					g.ExpectFail()
					g.Assert("Hello").Not.Eql("Hello")
				})

				g.It("IsTrue", func() {
					g.ExpectFail()
					g.Assert(true).Not.IsTrue()
				})

				g.It("IsFalse", func() {
					g.ExpectFail()
					g.Assert(false).Not.IsFalse()
				})

				g.It("IsOK", func() {
					g.ExpectFail()
					g.Assert(g).Not.IsOK()
				})

				g.It("IsType", func() {
					g.ExpectFail()
					g.Assert("").Not.HasSameType("")
				})
			})
		})
	})
}
