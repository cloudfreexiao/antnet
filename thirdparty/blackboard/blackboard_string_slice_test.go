package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardStringSlice(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when retrieving a non-existing StringSlice pointer with key `notexist`", func() {
			s := bb.StringSliceP("notexist")
			Convey("StringSlice pointer should be nil", func() {
				So(s, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetStringSlice with key `exist` and value `test1`, `test2`", func() {
			bb.SetStringSlice("exist", []string{"test1", "test2"})
			Convey("when retrieving an existing string pointer with key `exist`", func() {
				ss := bb.StringSliceP("exist")
				Convey("StringSlice pointer should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("string should equal `test1`, `test2`", func() {
					So((*ss)[0], ShouldEqual, "test1")
					So((*ss)[1], ShouldEqual, "test2")
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetStringSlice with key `exist` and value `test1`, `test2", func() {
			ssv := []string{"test1", "test2"}
			bb.SetStringSlice("exist", ssv)
			Convey("when retrieving an existing StringSlice pointer with key `exist`", func() {
				ss := bb.StringSliceP("exist")
				Convey("StringSlice pointer should not be nil", func() {
					So(ss, ShouldNotBeNil)
				})
				Convey("StringSlice should equal `test1`, `test2`", func() {
					So((*ss)[0], ShouldEqual, ssv[0])
					So((*ss)[1], ShouldEqual, ssv[1])
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetStringSliceP with key `exist` and value `test1`, `test2`", func() {
			ssv := []string{"test1", "test2"}
			bb.SetStringSliceP("exist", &ssv)
			Convey("when retrieving existing string slice pointers with AllStringSlice", func() {
				kssl := bb.AllStringSlice()
				Convey("KeyStringSlice slice should not be nil", func() {
					So(kssl, ShouldNotBeNil)
				})
				Convey("KeyStringSlice slice should have length of 1", func() {
					So(len(kssl), ShouldEqual, 1)
				})
				Convey("KeyStringSlice slice should contain {`exist`, correct string pointer}", func() {
					So(kssl, ShouldContain, KSS{"exist", &ssv})
				})
				Convey("StringSlice slice should equal `test1`, `test2`", func() {
					So((*kssl[0].Value), ShouldContain, ssv[0])
					So((*kssl[0].Value), ShouldContain, ssv[1])
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetValue with key `exist` and value `1`", func() {
			bb.SetValue("exist", 1)
			Convey("when retrieving value as StringSlice pointer with key `exist`", func() {
				s := bb.StringSliceP("exist")
				Convey("StringSlice pointer should be nil", func() {
					So(s, ShouldBeNil)
				})
			})
		})
	})
}
