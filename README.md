API Demo
========

This is a very simple demonstration of Golang library coding style which
prevents the caller from using unnamed struct initialization.  Doing this
ensures that new fields can safely be added to the struct without breaking
backwards compatibility.

Use `go run main.go` to see some output.

I've tried exploring various types, interfaces, fmt.Format and more, to try
to automatically hide the filtering field from `fmt.Printf` with `%#v` but to
no avail.  I ended up going for a very short name which idiomatically conveys
"unused" to hint at what's going on; could change the type away from
`struct{}` to `int` to shorten the `%#v` output even more.
