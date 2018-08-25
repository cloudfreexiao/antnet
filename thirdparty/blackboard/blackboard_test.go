package blackboard

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBlackboardNew(t *testing.T) {
	Convey("when givien a new blackboard", t, func() {
		bb := NewBlackboard()
		Convey("blackboard should not be nil", func() {
			So(bb, ShouldNotBeNil)
		})
	})
}

func TestBlackboardSingleton(t *testing.T) {
	Convey("when givien the blackboard singleton", t, func() {
		bb := Singleton()
		Convey("blackboard should not be nil", func() {
			So(bb, ShouldNotBeNil)
		})
	})
}
