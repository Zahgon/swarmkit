package deepcopy

// CopierFrom can be implemented if an object knows how to copy another into itself.
type CopierFrom interface {
	// Copy takes the fields from src and copies them into the target object.
	//
	// Calling this method with a nil receiver or a nil src may panic.
	CopyFrom(src interface{})
}

// Copy copies src into dst. dst and src must have the same type.
//
// If the type has a copy function defined, it will be used.
//
// Default implementations for builtin types and well known protobuf types may
// be provided.
//
// If the copy cannot be performed, this function will panic. Make sure to test
// types that use this function.
func Copy(dst, src interface{}) { _ = "STUB: not implemented"; return }
