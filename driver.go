//-----------------------------------------------------------------------------
// Package membership:
//-----------------------------------------------------------------------------

package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import (

	// Standard library:
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	// Community:
	"github.com/calavera/dkvolume"
)

//-----------------------------------------------------------------------------
// Structs definitions:
//-----------------------------------------------------------------------------

type rsyncDriver struct {
	srcdst                 map[string]string
	volRoot, sshKey        string
	archive, compress, del bool
}

//-----------------------------------------------------------------------------
// func rsyncArgs() returns a string of arguments to rsync
//-----------------------------------------------------------------------------

func (d *rsyncDriver) rsyncArgs() string {

	args := []string{}

	if d.archive {
		args = append(args, "--archive")
	}
	if d.compress {
		args = append(args, "--compress")
	}
	if d.del {
		args = append(args, "--delete")
	}

	// Remote shell customization:
	args = append(
		args,
		fmt.Sprintf(
			`-e 'ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o LogLevel=quiet -i "%s"'`,
			d.sshKey),
	)

	return strings.Join(args, " ")
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Create
//
// Request:
//  { "Name": "volume_name" }
//  Instruct the plugin that the user wants to create a volume, given a user
//  specified volume name. The plugin does not need to actually manifest the
//  volume on the filesystem yet (until Mount is called).
//
// Response:
//  { "Err": null }
//  Respond with a string error if an error occurred.
//-----------------------------------------------------------------------------

func (d *rsyncDriver) Create(r dkvolume.Request) dkvolume.Response {

	// Parse rsync source and destination:
	src := strings.Replace(r.Name, "/", ":/", 1) + "/"
	dst := filepath.Join(d.volRoot, r.Name) + "/"

	// Check whether remote is in use:
	if _, ok := d.srcdst[r.Name]; ok {
		log.Printf("Already in use: %s", src)
		return dkvolume.Response{Err: "Remote already in use"}
	}

	// Remember me:
	d.srcdst[r.Name] = dst

	// Create the destination directory:
	if err := os.MkdirAll(dst, 0755); err != nil {
		return dkvolume.Response{Err: err.Error()}
	}

	// Forge the command:
	command := "rsync " + d.rsyncArgs() + " " + src + " " + dst
	cmdRsync := exec.Command("/bin/sh", "-c", command)
	cmdRsync.Stdout = os.Stdout
	cmdRsync.Stderr = os.Stderr

	// Shellout rsync:
	log.Printf("Pulling data from: %s", src)
	if err := cmdRsync.Run(); err != nil {
		log.Printf("error: %v\n", err)
		return dkvolume.Response{Err: err.Error()}
	}

	// Return:
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Remove
//
// Request:
//  { "Name": "volume_name" }
//  Delete the specified volume from disk. This request is issued when a user
//  invokes docker rm -v to remove volumes associated with a container.
//
// Response:
//  { "Err": null }
//  Respond with a string error if an error occurred.
//-----------------------------------------------------------------------------

func (d *rsyncDriver) Remove(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Path
//
// Request:
//  { "Name": "volume_name" }
//  Docker needs reminding of the path to the volume on the host.
//
// Response:
//  { "Mountpoint": "/path/to/directory/on/host", "Err": null }
//  Respond with the path on the host filesystem where the volume has been
//  made available, and/or a string error if an error occurred.
//-----------------------------------------------------------------------------

func (d *rsyncDriver) Path(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{Mountpoint: d.srcdst[r.Name]}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Mount
//
// Request:
//  { "Name": "volume_name" }
//  Docker requires the plugin to provide a volume, given a user specified
//  volume name. This is called once per container start.
//
// Response:
//  { "Mountpoint": "/path/to/directory/on/host", "Err": null }
//  Respond with the path on the host filesystem where the volume has been
//  made available, and/or a string error if an error occurred.
//-----------------------------------------------------------------------------

func (d *rsyncDriver) Mount(r dkvolume.Request) dkvolume.Response {
	log.Printf("Mount point: %s\n", d.srcdst[r.Name])
	return dkvolume.Response{Mountpoint: d.srcdst[r.Name]}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Unmount
//
// Request:
//  { "Name": "volume_name" }
//  Indication that Docker no longer is using the named volume. This is called
//  once per container stop. Plugin may deduce that it is safe to deprovision
//  it at this point.
//
// Response:
//  { "Err": null }
//  Respond with a string error if an error occurred.
//-----------------------------------------------------------------------------

func (d *rsyncDriver) Unmount(r dkvolume.Request) dkvolume.Response {

	// Forge the command:
	src := strings.Replace(r.Name, "/", ":/", 1) + "/"
	command := "rsync " + d.rsyncArgs() + " " + d.srcdst[r.Name] + " " + src
	cmdRsync := exec.Command("/bin/sh", "-c", command)
	cmdRsync.Stdout = os.Stdout
	cmdRsync.Stderr = os.Stderr

	// Shellout rsync:
	log.Printf("Pushing data to %s\n", src)
	if err := cmdRsync.Run(); err != nil {
		log.Printf("error: %v\n", err)
		return dkvolume.Response{Err: err.Error()}
	}

	delete(d.srcdst, r.Name)
	return dkvolume.Response{}
}
