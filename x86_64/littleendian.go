// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// little endian

package x86_64

// uint64 little endian
func uint64le(v uint64) []uint8 {

	b := []uint8{} // le uint8 array

	b = append(
		b,
		uint8(v),
		uint8(v>>8),
		uint8(v>>16),
		uint8(v>>24),
		uint8(v>>32),
		uint8(v>>40),
		uint8(v>>48),
		uint8(v>>56),
	)

	return b
}

// uint32 little endian
func uint32le(v uint32) []uint8 {

	b := []uint8{} // le uint8 array

	b = append(
		b,
		uint8(v),
		uint8(v>>8),
		uint8(v>>16),
		uint8(v>>24),
	)

	return b
}

// uint16 little endian
func uint16le(v uint16) []uint8 {

	b := []uint8{} // le uint8 array

	b = append(
		b,
		uint8(v),
		uint8(v>>8),
	)

	return b
}

// uint8 little endian
func uint8le(v uint8) []uint8 {

	b := []uint8{} // le uint8 array

	b = append(
		b,
		uint8(v),
	)

	return b
}
