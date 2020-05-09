// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package badgerdb

import (
	"errors"
	"os"
	"time"

	"github.com/zerjioang/gotools/shared/db"

	"github.com/dgraph-io/badger/options"
	"github.com/zerjioang/gotools/lib/fastime"
	"github.com/zerjioang/gotools/util/codec"

	"github.com/zerjioang/gotools/lib/fs"

	"github.com/dgraph-io/badger"
	"github.com/zerjioang/gotools/lib/logger"
)

type BadgerStorage struct {
	// our custom database configuration options
	options *Options
	// instance is the underlying handle to the badgerdb.
	instance            *badger.DB
	vlogTicker          *time.Ticker // runs every 1m, check size of vlog and run GC conditionally.
	mandatoryVlogTicker *time.Ticker // runs every 10m, we always run vlog GC.
}

var (
	defaultConfig = badger.Options{
		LevelOneSize:        256 << 20,
		LevelSizeMultiplier: 10,
		TableLoadingMode:    options.MemoryMap, // options.LoadToRAM, // Mode in which LSM tree is loaded
		ValueLogLoadingMode: options.MemoryMap, // options.FileIO,    // options.MemoryMap,
		// table.MemoryMap to mmap() the tables.
		// table.Nothing to not preload the tables.
		MaxLevels:               7, // Size of table
		MaxTableSize:            64 << 20,
		NumCompactors:           1, // 2 // Number of concurrent compactions,
		NumLevelZeroTables:      5,
		NumLevelZeroTablesStall: 10,
		NumMemtables:            5,
		SyncWrites:              true,
		NumVersionsToKeep:       1,
		CompactL0OnClose:        true,
		// Nothing to read/write value log using standard File I/O
		// MemoryMap to mmap() the value log files
		// (2^30 - 1)*2 when mmapping < 2^31 - 1, max int32.
		// -1 so 2*ValueLogFileSize won't overflow on 32-bit systems.
		ValueLogFileSize: 1<<30 - 1,

		ValueLogMaxEntries: 1000000,
		ValueThreshold:     32,
		Truncate:           false,
		LogRotatesToFlush:  2,
	}
	uid = os.Getuid()
	gid = os.Getgid()
)

// database related errors
var (
	errDuplicateKey = errors.New("duplicate key found on database. cannot store")
)

func createData(path string) error {
	logger.Debug("creating badgerdb path")
	defaultConfig.Dir = path
	defaultConfig.ValueDir = path
	// create dir if does not exists
	if !fs.Exists(path) {
		logger.Debug("creating dir: ", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error("failed to create database dir:", err)
			return err
		}
	}
	// overwrite directory permissions
	logger.Debug("setting dir permissions and ownership: ", path)
	permErr := chownR(path, uid, gid, os.ModePerm)
	if permErr != nil {
		logger.Error("failed to set permissions: ", permErr)
	}
	return permErr
}

func NewCollection(databaseRootPath, name string) (*BadgerStorage, error) {
	logger.Debug("creating new badgerdb collection")
	err := createData(databaseRootPath + name)
	if err != nil {
		logger.Error("failed to create database dir", err)
		return nil, err
	}
	var openErr error
	collection := new(BadgerStorage)
	collection.instance, openErr = badger.Open(defaultConfig)
	if openErr != nil {
		logger.Error("failed to open badgerdb with default config", openErr)
		return nil, openErr
	}
	// register for listening poweroff events
	// todo integrate event bus
	/*
		bus.SharedBus().Subscribe(notifier.PowerOffEvent, func(message gobus.EventMessage) {
			logger.Debug("executing database poweroff routine in database: ", name)
			err := collection.Close()
			if err != nil {
				logger.Error("failed to close database due to: ", err)
			}
		})
	*/
	return collection, nil
}

func (storage *BadgerStorage) SetRawKey(key []byte, val []byte) error {
	logger.Debug("inserting key-value in badgerdb")
	return storage.instance.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (storage *BadgerStorage) SetCompositeKey(key []byte, val []byte) error {
	logger.Debug("inserting key-value in badgerdb")
	return storage.instance.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (storage *BadgerStorage) PutUniqueKeyValue(key []byte, value []byte) error {
	logger.Debug("inserting unique key-value in badgerdb")
	err := storage.instance.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == nil && item != nil {
			return errDuplicateKey
		} else {
			return txn.Set(key, value)
		}
	})
	return err
}

// Get is used to retrieve a value from the k/v store by key
func (storage *BadgerStorage) Get(key []byte) ([]byte, error) {
	logger.Debug("reading key from badgerdb")
	var value []byte
	err := storage.instance.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			switch err {
			case badger.ErrKeyNotFound:
				return badger.ErrKeyNotFound
			default:
				return err
			}
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (storage *BadgerStorage) Add(key, data []byte) error {
	logger.Debug("adding new key-value to badgerdb")
	return storage.SetRawKey(key, data)
}

func (storage *BadgerStorage) List(prefixStr string, decoderModel db.DaoInterface) ([]db.DaoInterface, error) {
	logger.Debug("listing values from badgerdb")
	var results []db.DaoInterface
	execErr := storage.instance.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		prefix := []byte(prefixStr)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			var readedVal []byte
			readedVal, err := item.ValueCopy(readedVal)
			if err != nil {
				it.Close()
				return err
			} else {
				//decode readed bytes to go struct
				readedModel, err := decoderModel.Decode(readedVal)
				if err.None() {
					results = append(results, readedModel)
				}
			}
		}
		it.Close()
		return nil
	})
	return results, execErr
}

func (storage *BadgerStorage) Delete(key []byte) error {
	logger.Debug("deleting key from badgerdb: ", string(key))
	execErr := storage.instance.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			logger.Error("failed to delete key from badgerdb: ", string(key), "error: ", err)
		}
		return err
	})
	return execErr
}

// DeleteRange deletes logs within a given range inclusively.
func (storage *BadgerStorage) DeleteRange(min, max uint64) error {
	logger.Debug("deleting by range from badgerdb")
	// we manage the transaction manually in order to avoid ErrTxnTooBig errors
	txn := storage.instance.NewTransaction(true)
	it := txn.NewIterator(badger.IteratorOptions{
		PrefetchValues: false,
		Reverse:        false,
	})

	seekKey := codec.Uint64ToBytes(min)
	for it.Seek(seekKey); it.Valid(); it.Next() {
		key := make([]byte, 8)
		it.Item().KeyCopy(key)
		//encode
		k := codec.BytesToUint64(key)
		// Handle out-of-range log index
		if k > max {
			break
		}
		// Delete in-range log index
		if err := txn.Delete(key); err != nil {
			if err == badger.ErrTxnTooBig {
				it.Close()
				err = txn.Commit()
				if err != nil {
					return err
				}
				return storage.DeleteRange(k, max)
			}
			return err
		}
	}
	it.Close()
	err := txn.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SetUint64 is like Set, but handles uint64 values
func (storage *BadgerStorage) SetUint64(key []byte, val uint64) error {
	return storage.SetRawKey(key, codec.Uint64ToBytes(val))
}

// GetUint64 is like Get, but handles uint64 values
func (storage *BadgerStorage) GetUint64(key []byte) (uint64, error) {
	val, err := storage.Get(key)
	if err != nil {
		return 0, err
	}
	return codec.BytesToUint64(val), nil
}

func (storage *BadgerStorage) runVlogGC(instance *badger.DB, threshold int64) {
	// Get initial size on start.
	_, lastVlogSize := instance.Size()

	runGC := func() {
		var err error
		for err == nil {
			// If a GC is successful, immediately run it again.
			err = instance.RunValueLogGC(0.7)
		}
		_, lastVlogSize = instance.Size()
	}

	for {
		select {
		case <-storage.vlogTicker.C:
			_, currentVlogSize := instance.Size()
			if currentVlogSize < lastVlogSize+threshold {
				continue
			}
			runGC()
		case <-storage.mandatoryVlogTicker.C:
			runGC()
		}
	}
}

// Close is used to gracefully close the DB connection.
func (storage *BadgerStorage) Close() error {
	logger.Debug("closing database")
	if storage.vlogTicker != nil {
		storage.vlogTicker.Stop()
	}
	if storage.mandatoryVlogTicker != nil {
		storage.mandatoryVlogTicker.Stop()
	}
	return storage.instance.Close()
}

// New uses the supplied options to open the Badger badgerdb and prepare it for
// use as a raft backend.
func NewBadgerStorageGC(options *Options) (*BadgerStorage, error) {

	// build badger options
	if options.BadgerOptions == nil {
		defaultOpts := badger.DefaultOptions("")
		options.BadgerOptions = &defaultOpts
	}
	options.BadgerOptions.Dir = options.Path
	options.BadgerOptions.ValueDir = options.Path
	options.BadgerOptions.SyncWrites = !options.NoSync

	// try to create new database handler
	storage, err := NewCollection(options.Path, "")
	if err != nil {
		return nil, err
	}
	storage.options = options

	// Start GC routine
	if options.ValueLogGC {

		var gcInterval fastime.Duration
		var mandatoryGCInterval fastime.Duration
		var threshold int64

		if gcInterval = 1 * fastime.Minute; options.GCInterval != 0 {
			gcInterval = options.GCInterval
		}
		if mandatoryGCInterval = 10 * fastime.Minute; options.MandatoryGCInterval != 0 {
			mandatoryGCInterval = options.MandatoryGCInterval
		}
		if threshold = int64(1 << 30); options.GCThreshold != 0 {
			threshold = options.GCThreshold
		}

		storage.vlogTicker = time.NewTicker(time.Duration(gcInterval))
		storage.mandatoryVlogTicker = time.NewTicker(time.Duration(mandatoryGCInterval))
		go storage.runVlogGC(storage.instance, threshold)
	}

	return storage, nil
}
