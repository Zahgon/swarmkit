package main

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// CustomNameID preprocess the field, and set the [(gogoproto.customname) = "..."]
// if necessary, in order to avoid setting `gogoproto.customname` manually.
// The automatically assigned name should conform to Golang convention.
func CustomNameID(file *descriptor.FileDescriptorProto) { _ = "STUB: not implemented"; return }

// Skip if [(gogoproto.customname) = "..."] has already been set.

// Skip if embedded

// id -> ID

// id_some -> IDSome

// some_id -> SomeID

// some_ids -> SomeIDs

// Iterate through all fields in file
