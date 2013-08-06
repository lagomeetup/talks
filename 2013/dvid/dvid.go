// This file has fragments of code used for the DVID presentation

// START_IMPORT OMIT
import (
    ...
	// Declare the data types this DVID executable will support
	_ "github.com/janelia-flyem/dvid/datatype/grayscale8"
	_ "github.com/janelia-flyem/dvid/datatype/labels32"
	_ "github.com/janelia-flyem/dvid/datatype/labels64"
	_ "github.com/janelia-flyem/dvid/datatype/rgba8"
	_ "github.com/DocSavage/dvid/datatype/raveler"
)
// END_IMPORT OMIT

// START_REGISTER OMIT
func init() {
	grayscale := voxels.NewDatatype()
	grayscale.DatatypeID = datastore.MakeDatatypeID("grayscale8", RepoUrl, Version)
	grayscale.NumChannels = 1
	grayscale.BytesPerVoxel = 1

	// Data types must be registered with the datastore to be used.
	datastore.RegisterDatatype(grayscale)
}
// END_REGISTER OMIT

const helpMessage = `
dvid is a distributed, versioned image-oriented datastore

Usage: dvid [options] <command>

      -datastore  =string   Path to DVID datastore directory (default: current directory).
      -webclient  =string   Path to web client directory.  Leave unset for default pages.
      -rpc        =string   Address for RPC communication.
      -http       =string   Address for HTTP communication.
      -cpuprofile =string   Write CPU profile to this file.
      -memprofile =string   Write memory profile to this file on ctrl-C.
      -numcpu     =number   Number of logical CPUs to use for DVID.
      -timeout    =number   Seconds to wait trying to get exclusive access to datastore.
      -gzip       (flag)    Turn gzip compression on for REST API.
      -types      (flag)    Show compiled DVID data types
      -debug      (flag)    Run in debug mode.  Verbose.
      -benchmark  (flag)    Run in benchmarking mode. 
  -h, -help       (flag)    Show help message

  For profiling, please refer to this excellent article:
  http://blog.golang.org/2011/06/profiling-go-programs.html

Commands that can be performed without a running server:

	about
	help
	init 
	serve
`

// START_STORAGE OMIT
// KeyValueDB provides an interface to a number of key/value
// implementations, e.g., C++ or Go leveldb libraries.
type KeyValueDB interface {
	// Closes datastore.
	Close()

	// Get returns a value given a key.
	Get(k Key) (v []byte, err error)

	// Put writes a value with given key.
	Put(k Key, v []byte) (err error)

	// Delete removes an entry given key
	Delete(k Key) (err error)
}
// END_STORAGE OMIT

// START_CONDITIONAL OMIT
// +build levigo // HL

package storage

import (
	"log"

	"github.com/janelia-flyem/dvid/dvid"
	"github.com/janelia-flyem/go/levigo" // HL
)
// END_CONDITIONAL OMIT

// NewIndexIterator returns an IndexIterator closure.  The closure for this voxels data
// type uses VoxelCoord, BlockCoord, and returns an IndexZYX.
// START_ITERATOR OMIT
func (dset *Dataset) NewIndexIterator(extents interface{}) datastore.IndexIterator {
	data, ok := extents.(dvid.Geometry)
	...
	x, y, z := startBlockCoord[0], startBlockCoord[1], startBlockCoord[2]
	...
	return func() datastore.Index { // HL
		if z > endBlockCoord[2] {
			return nil
		}
		// Convert an n-D block coordinate into a 1-D index
		index, err := datastore.BlockZYX{dvid.BlockCoord{x, y, z}}.MakeIndex()
		x++
		if x > endBlockCoord[0] {
			x = startBlockCoord[0]
			y++
		}
		if y > endBlockCoord[1] {
			y = startBlockCoord[1]
			z++
		}
		return index
	}
}
// END_ITERATOR OMIT

// START_DATATYPE OMIT
type TypeID interface {
	DatatypeName() string

	DatatypeUrl() UrlString

	DatatypeVersion() string
}

type TypeService interface {
	TypeID

	Help() string

	NewDataset(id DatasetID, config dvid.Config) DatasetService
}
// END_DATATYPE OMIT

// START_DATASET OMIT
type DatasetService interface {
	TypeService

	// DatasetName returns the name of the dataset.
	DatasetName() DatasetString

	// ChunkHandler processes a chunk of data in a mapped operation.
	ChunkHandler(op *ChunkOp) // HL

	// Handle iteration through a dataset in abstract way.
	NewIndexIterator(extents interface{}) IndexIterator

	// DoRPC handles command line and RPC commands specific to a data type
	DoRPC(request Request, reply *Response) error

	// DoHTTP handles HTTP requests specific to a data type
	DoHTTP(w http.ResponseWriter, r *http.Request) error

	// Returns standard error response for unknown commands
	UnknownCommand(r Request) error
}
// END_DATASET OMIT

// START_GET OMIT
func (dset *Dataset) GetImage(versionID dvid.LocalID, slice dvid.Geometry) 
    (img image.Image, err error) {
        
	bytesPerVoxel := dset.BytesPerVoxel()
	stride := slice.Width() * bytesPerVoxel

	numBytes := int64(bytesPerVoxel) * slice.NumVoxels()
	data := make([]uint8, numBytes, numBytes)
	operation := dset.makeOp(&Voxels{slice, data, stride}, versionID, GetOp)

	// Perform operation using mapping
	err = operation.Map() // HL
	if err != nil {
		return
	}
	operation.Wait() // HL

	img, err = dset.SliceImage(operation.data, 0)
	return
}
// END_GET OMIT

// START_MAP OMIT
func Map(op Operation) error {
	channels, found := chunkChannels[op.DatasetName()]
	indexIterator := op.IndexIterator() // HL

	DiskAccess.Lock()
	for {
		index := indexIterator()
		if index == nil {
			break
		}
		key := storage.Key{datasetID, versionID, index.Bytes()}

		// Pull from the datastore ...

		op.WaitAdd() // HL
		channelNum := index.Hash(len(channels))
		channels[channelNum] <- &ChunkOp{op, value, key, false} // HL
	}
	DiskAccess.Unlock()
	return nil
}
// END_MAP OMIT
