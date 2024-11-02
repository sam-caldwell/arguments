package types

import (
	"fmt"
	"strconv"
	"unsafe"
)

// Uint16 - represents an unsigned 16-bit integer type.
type Uint16 uint16

// String - converts the Uint16 to a string.
func (u *Uint16) String() string {
	return strconv.FormatUint(uint64(*u), 10)
}

// Parse - converts a string to a Uint16 value.
func (u *Uint16) Parse(v string) error {
	sz := int(unsafe.Sizeof(uint16(0))) * 8
	n, err := strconv.ParseUint(v, 10, sz)
	if err != nil {
		return fmt.Errorf("parse error: %v (input: %v, sz: %d)", err, v, sz)
	}
	*u = Uint16(n)
	return nil
}
