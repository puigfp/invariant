# invariant

This short package provides helpers for checking invariant in your Go programs.

Using this package, you can easily write debug assertions (ignored in "release" mode) and regular assertions, that will make your program panic if they are violated.

Please read the [GoDoc](https://godoc.org/github.com/puigfp/invariant) for more information.

## Example

I'm sure that your real world programs have more interesting invariants to check than this one.

```go
// Fibonnacci returns what you think it returns
// n should be a positive integer.
func Fibonacci(n int) int64 {
    invariant.Assert(n >= 0, "n should be positive")

    var (
        a int64 = 0
        b int64 = 1
    )

    for n > 0 {
        a, b = b, a + b
        n--
    }

    return a
}
```

## References

This is inspired from an article by Matt Klein: [Crash early and crash often for more reliable software](https://medium.com/@mattklein123/crash-early-and-crash-often-for-more-reliable-software-597738dd21c5)


