package badgerdb

import (
	"github.com/dgraph-io/badger"
	"github.com/zerjioang/gotools/lib/fastime"
)

// Options contains all the configuration used to open the Badger badgerdb
type Options struct {
	// Path is the directory path to the Badger badgerdb to use.
	Path string

	// BadgerOptions contains any specific Badger options you might
	// want to specify.
	BadgerOptions *badger.Options

	// NoSync causes the database to skip fsync calls after each
	// write to the log. This is unsafe, so it should be used
	// with caution.
	NoSync bool

	// ValueLogGC enables a periodic goroutine that does a garbage
	// collection of the value log while the underlying Badger is online.
	ValueLogGC bool

	// GCInterval is the interval between conditionally running the garbage
	// collection process, based on the size of the vlog. By default, runs every 1m.
	GCInterval fastime.Duration

	// GCInterval is the interval between mandatory running the garbage
	// collection process. By default, runs every 10m.
	MandatoryGCInterval fastime.Duration

	// GCThreshold sets threshold in bytes for the vlog size to be included in the
	// garbage collection cycle. By default, 1GB.
	GCThreshold int64
}
