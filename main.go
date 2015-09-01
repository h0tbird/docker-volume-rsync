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
	"os/user"
	"path/filepath"

	// Community:
	"github.com/calavera/dkvolume"
)

//-----------------------------------------------------------------------------
// Package constant declarations factored into a block:
//-----------------------------------------------------------------------------

const (
	id            = "rsync"
	socketAddress = "/var/run/docker/plugins/rsync.sock"
)

//-----------------------------------------------------------------------------
// Package variable declarations factored into a block:
//-----------------------------------------------------------------------------

var (

	// Predefined defaults:
	defVolRoot = filepath.Join(dkvolume.DefaultDockerRootDirectory, id)
	defSshKey  = sshKeyPath()

	// Flags:
	volRoot  = flag.String("volroot", defVolRoot, "Docker volumes root directory")
	archive  = flag.Bool("archive", true, "Archive mode; equals -rlptgoD")
	del      = flag.Bool("delete", false, "Delete extraneous files from dest dirs")
	compress = flag.Bool("compress", false, "Compress file data during the transfer")
	sshKey   = flag.String("ssh-private-key", defSshKey, "Path to the private SSH key")
)

//-----------------------------------------------------------------------------
// func sshKeyPath() returns the path to the current user's default private
// ssh key file.
//-----------------------------------------------------------------------------

func sshKeyPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".ssh/id_rsa")
}

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
	d := rsyncDriver{
		srcdst:   make(map[string]string),
		volRoot:  *volRoot,
		sshKey:   *sshKey,
		archive:  *archive,
		del:      *del,
		compress: *compress,
	}

	// Initializes the request handler with a driver implementation:
	h := dkvolume.NewHandler(&d)

	// Listen for requests in a unix socket:
	log.Printf("Listening on %s\n", socketAddress)
	fmt.Println(h.ServeUnix("root", socketAddress))
}
