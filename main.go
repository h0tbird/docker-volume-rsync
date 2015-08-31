//-----------------------------------------------------------------------------
// This volume driver is meant to be used by docker >= 1.8.x
//
// 1- run the driver:
// sudo docker-volume-rsync
//
// 2- run the container:
// docker run -it --volume-driver rsync -v src.host.org/foo:/foo alpine sh
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Package membership:
//-----------------------------------------------------------------------------

package main

//-----------------------------------------------------------------------------
// Package factored import statement:
//-----------------------------------------------------------------------------

import (

	// Standard library:
	"flag"
	"fmt"
	"log"
	"os"

	// Community:
	"github.com/calavera/dkvolume"
)

//-----------------------------------------------------------------------------
// Package constant declarations factored into a block:
//-----------------------------------------------------------------------------

const (
	socketAddress = "/var/run/docker/plugins/rsync.sock"
)

//-----------------------------------------------------------------------------
// Package variable declarations factored into a block:
//-----------------------------------------------------------------------------

var (
	archive  = flag.Bool("archive", true, "archive mode; equals -rlptgoD")
	compress = flag.Bool("compress", false, "compress file data during the transfer")
)

//-----------------------------------------------------------------------------
// func init() is called after all the variable declarations in the package
// have evaluated their initializers, and those are evaluated only after all
// the imported packages have been initialized:
//-----------------------------------------------------------------------------

func init() {

	// Check for mandatory argc:
	if len(os.Args) < 1 {
		usage()
	}

	// Parse commandline flags:
	flag.Usage = usage
	flag.Parse()
}

//-----------------------------------------------------------------------------
// func usage() reports the correct commandline usage for this driver:
//-----------------------------------------------------------------------------

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

//-----------------------------------------------------------------------------
// Function main of package main:
//-----------------------------------------------------------------------------

func main() {

	// Initialize the driver struct:
	d := rsyncDriver{}

	// Initializes the request handler with a driver implementation:
	h := dkvolume.NewHandler(d)

	// Listen for requests in a unix socket:
	log.Printf("Listening on %s\n", socketAddress)
	fmt.Println(h.ServeUnix("root", socketAddress))
}
