package gorocksdb

// #include "rocksdb/c.h"
import "C"

type EnvOptions struct {
	c *C.rocksdb_envoptions_t
}

// NewDefaultIngestExternalFileOptions creates a default IngestExternalFileOptions object.
func NewDefaultEnvOptions() *EnvOptions {
	return NewNativeEnvOptions(C.rocksdb_envoptions_create())
}

// NewNativeFlushOptions creates a FlushOptions object.
func NewNativeEnvOptions(c *C.rocksdb_envoptions_t) *EnvOptions {
	return &EnvOptions{c: c}
}
