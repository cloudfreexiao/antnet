package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardBool(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when retrieving a non-existing Bool pointer with key `notexist`", func() {
			b := bb.BoolP("notexist")
			Convey("Bool pointer should be nil", func() {
				So(b, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetBool with key `exist` and value `true`", func() {
			bb.SetBool("exist", true)
			Convey("when retrieving an existing Bool pointer with key `exist`", func() {
				b := bb.BoolP("exist")
				Convey("Bool pointer should not be nil", func() {
					So(b, ShouldNotBeNil)
				})
				Convey("Bool should be `true`", func() {
					So(*b, ShouldBeTrue)
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetBoolP with key `exist` and value `true`", func() {
			v := true
			bb.SetBoolP("exist", &v)
			Convey("when retrieving an existing Bool pointer with key `exist`", func() {
				b := bb.BoolP("exist")
				Convey("Bool pointer should not be nil", func() {
					So(b, ShouldNotBeNil)
				})
				Convey("Bool should be `true`", func() {
					So(*b, ShouldBeTrue)
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetBoolP with key `exist` and value `true`", func() {
			v := true
			bb.SetBoolP("exist", &v)
			Convey("when retrieving existing Bool pointers with AllBool", func() {
				kbl := bb.AllBool()
				Convey("KeyBool slice should not be nil", func() {
					So(kbl, ShouldNotBeNil)
				})
				Convey("KeyBool slice should have length of 1", func() {
					So(len(kbl), ShouldEqual, 1)
				})
				Convey("KeyBool slice should contain {`exist`, correct Bool pointer}", func() {
					So(kbl, ShouldContain, KB{"exist", &v})
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetValue with key `exist` and value `1`", func() {
			bb.SetValue("exist", 1)
			Convey("when retrieving value as Bool pointer with key `exist`", func() {
				b := bb.BoolP("exist")
				Convey("Bool pointer should be nil", func() {
					So(b, ShouldBeNil)
				})
			})
		})
	})
}
