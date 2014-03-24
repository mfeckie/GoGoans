package goans

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type FillMeInToPassTheTest struct {
	message string
}

var __ = FillMeInToPassTheTest{message: "Fill me in to pass the test"}

func TestStrings(t *testing.T) {

	Convey("Learning about strings in Go", t, func() {
		Convey("Strings are Strings!", func() {
			testString := "Welcome to Go!"
			So(__, ShouldHaveSameTypeAs, testString) //Is a number the same as a string?
		})

		Convey("Strings can be created with `backticks` too", func() {
			backTickString := `Welcome to Go!`
			So(__, ShouldHaveSameTypeAs, backTickString)
		})

		Convey("Strings can be joined", func() {
			firstString := "Welcome to "
			secondString := "Go!"
			So(__, ShouldEqual, firstString+secondString) //What should our end result be?
			Convey("Joining should not change the original", func() {
				So(__, ShouldEqual, firstString)
				So(__, ShouldEqual, secondString)
			})
		})

		Convey("Some characters need to be escaped in string e.g.( \" )", func() {
			escaped := "Hello \"World!\""
			So(__, ShouldEqual, escaped)
		})

		Convey("Strings can be accessed like arrays, but don't return letters!", func() {
			stringAccessedLikeArray := "Hello Go!"
			So(__, ShouldHaveSameTypeAs, stringAccessedLikeArray[1]) // byte(123) is a way of creating a uint8
		})

		Convey("We can update a string by adding to it", func() {
			stringOne := "Hello "
			stringTwo := "Go!"
			stringOne += stringTwo
			So(__, ShouldEqual, "Hello Go!")
			Convey(",but this is a destructive operation", func() {
				So(__, ShouldNotEqual, "Hello ")
			})
		})
	})
}
