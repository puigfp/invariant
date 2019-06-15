package invariant

import "testing"

// assertPanic is a helper that asserts that a test panics
func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Fail()
	}
}

// overrideModeHelper is a helper used to override `mode`'s value (package variable) during a single test
func overrideModeHelper(tmpMode string) func() {
	backupMode := mode
	mode = tmpMode
	return func() {
		mode = backupMode
	}
}

func TestAssert(t *testing.T) {

	t.Run("invariant violated", func(t *testing.T) {
		defer assertPanic(t)

		Assert(false, "msg")
	})

	t.Run("invariant holds", func(t *testing.T) {
		Assert(true, "msg")
	})
}

func TestAssertDebug(t *testing.T) {

	t.Run("debug mode, invariant violated", func(t *testing.T) {
		defer assertPanic(t)

		revert := overrideModeHelper("debug")
		defer revert()

		AssertDebug(false, "msg")
	})

	t.Run("release mode, invariant violated", func(t *testing.T) {
		revert := overrideModeHelper("release")
		defer revert()

		AssertDebug(false, "msg")
	})

	t.Run("invariant holds", func(t *testing.T) {
		AssertDebug(true, "msg")
	})
}
