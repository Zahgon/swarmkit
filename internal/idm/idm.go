// Package idm manages reservation/release of numerical ids from a configured set of contiguous ids.
package idm

import (
	"errors"

	"github.com/bits-and-blooms/bitset"
)

var (
	// ErrNoBitAvailable is returned when no more bits are available to set
	ErrNoBitAvailable = errors.New("no bit available")
	// ErrBitAllocated is returned when the specific bit requested is already set
	ErrBitAllocated = errors.New("requested bit is already allocated")
)

// IDM manages the reservation/release of numerical ids from a contiguous set.
//
// An IDM instance is not safe for concurrent use.
type IDM struct {
	start, end uint
	set        *bitset.BitSet
	next       uint // index of the bit to start searching for the next serial allocation from (not offset by start)
}

// New returns an instance of id manager for a [start,end] set of numerical ids.
func New(start, end uint) (*IDM, error) { _ = "STUB: not implemented"; return nil, nil }

// GetID returns the first available id in the set.
func (i *IDM) GetID(serial bool) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

// GetSpecificID tries to reserve the specified id.
func (i *IDM) GetSpecificID(id uint) error { _ = "STUB: not implemented"; return nil }

// Release releases the specified id.
func (i *IDM) Release(id uint) { _ = "STUB: not implemented"; return }
