package foo

import (
	"fmt"
)

type Foo struct {
	Alpha string
	Beta  int

	// force the caller to use named fields from outside
	_ struct{}
}

// XGoString might be renamed GoString if you want to explore a little.
//
// GoString hides the coercion of forcing named fields from outside from fmt.*printf
// This is equivalent to the protection against encoding formats via reflect hiding
// unexported symbols, but applying to fmt.
// Downsides are:
// * we have to list every field we do export
// * we lose sight of whether the variable is a pointer or not
//   (a leading '&' which might be in %#v output is not shown)
// Upsides are:
// * we hide ugliness from the caller
//
// Unfortunately, there's no way to use `fmt:"-"` on the member, and the fmt
// printReflectValue() has no way to suppress the item; I've looked through
// various interfaces available, and even using a custom Format(fmt.State, rune)
// does not let us hide the element.  So we can either give it an explicit name
// which we have to hide like this, or give is an '_' name so it's as
// unobnoxious as possible when printed, so that if you do not enable this method
// then you can still get fairly readable data.  (use a string or int
// to shorten the output even more, instead of struct{})
func (f Foo) XGoString() string {
	return fmt.Sprintf("%T{Alpha:%q, Beta:%d}", f, f.Alpha, f.Beta)
}
