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
	flagA = flag.Bool("flagA", false, "Flag A is true or false")
	flagB = flag.Bool("flagB", false, "Flag B is true or false")
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
