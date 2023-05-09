// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// architecture modrm table

package x86_64

// modrm byte table
var archModrmList map[*archOperand][]uint8 = map[*archOperand][]uint8{
	{t: "", v: 0}: []uint8{},
}
