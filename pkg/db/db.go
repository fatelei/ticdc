// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package db

// DB is an interface of a leveldb-like database.
type DB interface {
	Snapshot() (Snapshot, error)
	Batch(cap int) Batch
	Close() error
	CollectMetrics(captureAddr string, id int)
}

// A Batch is a sequence of Puts and Deletes that Commit to DB.
type Batch interface {
	Put(key, value []byte)
	Delete(key []byte)
	Commit() error
	Count() uint32
	Repr() []byte
	Reset()
}

// Snapshot is an interface of a point-in-time view of the current DB state.
type Snapshot interface {
	Iterator(lowerBound, upperBound []byte) Iterator
	Release() error
}

// Iterator is an interface of an iterator of a DB.
type Iterator interface {
	Valid() bool
	First() bool
	Seek([]byte) bool
	Next() bool
	Key() []byte
	Value() []byte
	Error() error
	Release() error
}
