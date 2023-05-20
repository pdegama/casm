// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// little endian

package x86_64

// uint64 little endian
func uint64le(v uint64) []byte {

	b := []byte{} // le byte array

	b = append(
		b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
		byte(v>>32),
		byte(v>>40),
		byte(v>>48),
		byte(v>>56),
	)

	return b
}

// uint32 little endian
func uint32le(v uint32) []byte {

	b := []byte{} // le byte array

	b = append(
		b,
		byte(v),
		byte(v>>8),
		byte(v>>16),
		byte(v>>24),
	)

	return b
}

// uint16 little endian
func uint16le(v uint16) []byte {

	b := []byte{} // le byte array

	b = append(
		b,
		byte(v),
		byte(v>>8),
	)

	return b
}

// uint8 little endian
func uint8le(v uint8) []byte {

	b := []byte{} // le byte array

	b = append(
		b,
		byte(v),
	)

	return b
}
