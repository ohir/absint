// package absint provides abs() functions for the integers

package absint

// y := signbit; return (x ⨁  y) - y
func abs64(n int64) int64   { y := n >> 63; return (n ^ y) - y }
func abs32(n int32) int32   { y := n >> 31; return (n ^ y) - y }
func abs16(n int16) int16   { y := n >> 15; return (n ^ y) - y }
func abs8(n int8) int8      { y := n >> 7; return (n ^ y) - y }
func abs64u(n int64) uint64 { y := n >> 63; ; return uint64((n ^ y) - y) }
func abs32u(n int32) uint32 { y := n >> 31; return uint32((n ^ y) - y) }
func abs16u(n int16) uint16 { y := n >> 15; return uint16((n ^ y) - y) }
func abs8u(n int8) uint8    { y := n >> 7; return uint8((n ^ y) - y) }
