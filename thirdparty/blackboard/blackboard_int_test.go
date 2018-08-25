package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardInt(t *testing.T) {

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when retrieving a non-existing Int pointer with key `notexist`", func() {
			i := bb.IntP("notexist")
			Convey("Int pointer should be nil", func() {
				So(i, ShouldBeNil)
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetInt with key `exist` and value `10`", func() {
			bb.SetInt("exist", 10)
			Convey("when retrieving an existing Int pointer with key `exist`", func() {
				i := bb.IntP("exist")
				Convey("Int pointer should not be nil", func() {
					So(i, ShouldNotBeNil)
				})
				Convey("Int should equal `10`", func() {
					So(*i, ShouldEqual, 10)
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetIntP with key `exist` and value `10`", func() {
			v := 10
			bb.SetIntP("exist", &v)
			Convey("when retrieving an existing Int pointer with key `exist`", func() {
				i := bb.IntP("exist")
				Convey("Int pointer should not be nil", func() {
					So(i, ShouldNotBeNil)
				})
				Convey("Int should equal `10`", func() {
					So(*i, ShouldEqual, 10)
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetIntP with key `exist` and value `10`", func() {
			v := 10
			bb.SetIntP("exist", &v)
			Convey("when retrieving existing Int pointers with AllInt", func() {
				kil := bb.AllInt()
				Convey("KeyInt slice should not be nil", func() {
					So(kil, ShouldNotBeNil)
				})
				Convey("KeyInt slice should have length of 1", func() {
					So(len(kil), ShouldEqual, 1)
				})
				Convey("KeyInt slice should contain {`exist`, correct int pointer}", func() {
					So(kil, ShouldContain, KI{"exist", &v})
				})
			})
		})
	})

	Convey("when given a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("when using SetValue with key `exist` and value `true`", func() {
			bb.SetValue("exist", true)
			Convey("when retrieving value as Int pointer with key `exist`", func() {
				i := bb.IntP("exist")
				Convey("String pointer should be nil", func() {
					So(i, ShouldBeNil)
				})
			})
		})
	})
}
