// Package invariant holds short helpers you can use to check invariants.
//
// By default, this package runs in debug node, which means that all your assertions will be run.
//
// You can change the mode to "release"at compile time:
//
// 	 	go build -ldflags “-X github.com/puigfp/invariant.mode=release"
//
// In release mode, debug assertions are skipped.
//
// mode should be equal to "debug" or "release" (if this isn't true, the program will panic at
// startup).
package invariant
