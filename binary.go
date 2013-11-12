// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binary implements various bit- and byte-level
// manipulations.
package binary

// Reverse the bits in the binary representation
// of u. For example, 1011 becomes 1101.
func ReverseBits8(u uint8) uint8 {
	u = ((u & 0x0F) << 4) | ((u & 0xF0) >> 4)
	u = ((u & 0x33) << 2) | ((u & 0xCC) >> 2)
	return ((u & 0x55) << 1) | ((u & 0xAA) >> 1)
}

// Reverse the bits in the binary representation
// of u. For example, 1011 becomes 1101.
func ReverseBits16(u uint16) uint16 {
	u = ((u & 0x00FF) << 8) | ((u & 0xFF00) >> 8)
	u = ((u & 0x0F0F) << 4) | ((u & 0xF0F0) >> 4)
	u = ((u & 0x3333) << 2) | ((u & 0xCCCC) >> 2)
	return ((u & 0x5555) << 1) | ((u & 0xAAAA) >> 1)
}

// Reverse the bits in the binary representation
// of u. For example, 1011 becomes 1101.
func ReverseBits32(u uint32) uint32 {
	u = ((u & 0x0000FFFF) << 16) | ((u & 0xFFFF0000) >> 16)
	u = ((u & 0x00FF00FF) << 8) | ((u & 0xFF00FF00) >> 8)
	u = ((u & 0x0F0F0F0F) << 4) | ((u & 0xF0F0F0F0) >> 4)
	u = ((u & 0x33333333) << 2) | ((u & 0xCCCCCCCC) >> 2)
	return ((u & 0x55555555) << 1) | ((u & 0xAAAAAAAA) >> 1)
}

// Reverse the bits in the binary representation
// of u. For example, 1011 becomes 1101.
func ReverseBits64(u uint64) uint64 {
	u = ((u & 0x00000000FFFFFFFF) << 32) | ((u & 0xFFFFFFFF00000000) >> 32)
	u = ((u & 0x0000FFFF0000FFFF) << 16) | ((u & 0xFFFF0000FFFF0000) >> 16)
	u = ((u & 0x00FF00FF00FF00FF) << 8) | ((u & 0xFF00FF00FF00FF00) >> 8)
	u = ((u & 0x0F0F0F0F0F0F0F0F) << 4) | ((u & 0xF0F0F0F0F0F0F0F0) >> 4)
	u = ((u & 0x3333333333333333) << 2) | ((u & 0xCCCCCCCCCCCCCCCC) >> 2)
	return ((u & 0x5555555555555555) << 1) | ((u & 0xAAAAAAAAAAAAAAAA) >> 1)
}
