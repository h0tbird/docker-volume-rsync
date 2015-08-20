package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import (

	// External imports:
	"github.com/calavera/dkvolume"
)

//-----------------------------------------------------------------------------
// Structures definitions.
//-----------------------------------------------------------------------------

type myDummyDriver struct {
	name string
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

func (d myDummyDriver) Create(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Remove
//-----------------------------------------------------------------------------

func (d myDummyDriver) Remove(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Path
//-----------------------------------------------------------------------------

func (d myDummyDriver) Path(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Mount
//-----------------------------------------------------------------------------

func (d myDummyDriver) Mount(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}

//-----------------------------------------------------------------------------
// POST /VolumeDriver.Unmount
//-----------------------------------------------------------------------------

func (d myDummyDriver) Unmount(r dkvolume.Request) dkvolume.Response {
	return dkvolume.Response{}
}
