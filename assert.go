package invariant

import "fmt"

// mode is the assertion level this package will use.
var mode = "debug"

func init() {
	// make sure `mode` wasn't overriden with a bad value
	Assert(
		mode == "debug" || mode == "release",
		fmt.Sprintf("github.com/puigfp/invariant.mode should equal \"debug\" or "+
			"\"release\", value is: \"%s\"", mode),
	)
}

// Assert asserts the passed invariant, `panic(v)` is called if that invariant is violated.
func Assert(invariant bool, v interface{}) {
	if !invariant {
		panic(v)
	}
}

// AssertDebug behaves like Assert, but it skips the check if package is in release mode.
func AssertDebug(invariant bool, v interface{}) {
	if mode != "debug" {
		return
	}

	Assert(invariant, v)
}
