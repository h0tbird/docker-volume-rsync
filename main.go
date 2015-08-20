package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Native imports:
	"flag"
	"fmt"

	// External imports:
	"github.com/calavera/dkvolume"
)

//-----------------------------------------------------------------------------
// Package variable declarations factored into a block:
//-----------------------------------------------------------------------------

var (
	flag_a = flag.Bool("flag_a", false, "Flag a is true or false")
	flag_b = flag.Bool("flag_b", false, "Flag b is true or false")
)

//-----------------------------------------------------------------------------
// func init() is called after all the variable declarations in the package
// have evaluated their initializers, and those are evaluated only after all
// the imported packages have been initialized:
//-----------------------------------------------------------------------------

func init() {
	flag.Parse()
}

//-----------------------------------------------------------------------------
// Function main of package main:
//-----------------------------------------------------------------------------

func main() {

	fmt.Printf("Hello World!\n")
	d := myDummyDriver{}
	h := dkvolume.NewHandler(d)
	h.ServeUnix("root", "dummy_volume")
}
