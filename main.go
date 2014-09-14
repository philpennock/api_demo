package main

import (
	"fmt"

	"github.com/philpennock/api_demo/foo"
)

func main() {
	// This works:
	f := foo.Foo{Alpha: "text", Beta: 3}

	// This would yield: implicit assignment of unexported field 'forceNamedFieldsFromOutside' in foo.Foo literal
	//x := struct{}{}
	//f := foo.Foo{"text", 3, x}

	ff := &f
	fmt.Printf(" f: %T %#v\n", f, f)
	fmt.Printf("ff: %T %#v\n", ff, ff)
}
