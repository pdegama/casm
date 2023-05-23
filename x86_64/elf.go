// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// elf

package x86_64

import (
	"fmt"
	"os"
)

// elf structure
type elf struct {
	elfBytes []uint8
}

const (
	virtualStartAddress     uint64 = 0x400000
	dataVirtualStartAddress uint64 = 0x600000
	alignment               uint64 = 0x200000
)

// build elf
func (e *elf) buildELF(b *binaryGen) {

	bin := b.getBin()

	textSize := uint64(len(bin))
	// Size of ELF header + 2 * size program header?
	textOffset := uint64(0x40 + (2 * 0x38))

	e.addBytes(0x7f, 0x45, 0x4c, 0x46) // ELF magic value

	e.addBytes(0x02) // 64-bit executable
	e.addBytes(0x01) // Little endian
	e.addBytes(0x01) // ELF version
	e.addBytes(0x00) // Target OS ABI
	e.addBytes(0x00) // Further specify ABI version

	e.addBytes(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00) // Unused bytes

	e.addBytes(0x02, 0x00)             // Executable type
	e.addBytes(0x3e, 0x00)             // x86-64 target architecture
	e.addBytes(0x01, 0x00, 0x00, 0x00) // ELF version

	// 64-bit virtual offsets always start at 0x400000?? https://stackoverflow.com/questions/38549972/why-elf-executables-have-a-fixed-load-address
	// This seems to be a convention set in the x86_64 system-v abi: https://refspecs.linuxfoundation.org/elf/x86_64-SysV-psABI.pdf P26
	e.addValue(8, virtualStartAddress+textOffset)

	e.addBytes(0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00) // Offset from file to program header
	e.addBytes(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00) // Start of section header table
	e.addBytes(0x00, 0x00, 0x00, 0x00)                         // Flags
	e.addBytes(0x40, 0x00)                                     // Size of this header
	e.addBytes(0x38, 0x00)                                     // Size of a program header table entry - This should always be the same for 64-bit
	e.addBytes(0x02, 0x00)                                     // Length of sections: data and text for now
	e.addBytes(0x00, 0x00)                                     // Size of section header, which we aren't using
	e.addBytes(0x00, 0x00)                                     // Number of entries section header
	e.addBytes(0x00, 0x00)                                     // Index of section header table entry

	// Build Program Header
	// Text Segment
	e.addBytes(0x01, 0x00, 0x00, 0x00) // PT_LOAD, loadable segment. Both data and text segment use this.
	e.addBytes(0x07, 0x00, 0x00, 0x00) // Flags: 0x4 executable, 0x2 write, 0x1 read
	e.addValue(8, 0)                   // textOffset)          // Offset from the beginning of the file. These values depend on how big the header and segment sizes are.
	e.addValue(8, virtualStartAddress)
	e.addValue(8, virtualStartAddress) // Physical address, irrelavnt on linux.
	e.addValue(8, textSize)            // Number of bytes in file image of segment, must be larger than or equal to the size of payload in segment. Should be zero for bss data.
	e.addValue(8, textSize)            // Number of bytes in memory image of segment, is not always same size as file image.
	e.addValue(8, alignment)

	dataSize := uint64(len([]uint8{}))
	dataOffset := uint64(textOffset + textSize)
	dataVirtualAddress := dataVirtualStartAddress + dataOffset

	// Build Program Header
	// Data Segment
	e.addBytes(0x01, 0x00, 0x00, 0x00) // PT_LOAD, loadable segment. Both data and text segment use this.
	e.addBytes(0x07, 0x00, 0x00, 0x00) // Flags: 0x4 executable, 0x2 write, 0x1 read
	e.addValue(8, dataOffset)          // Offset address.
	e.addValue(8, dataVirtualAddress)  // Virtual address.
	e.addValue(8, dataVirtualAddress)  // Physical address.
	e.addValue(8, dataSize)            // Number of bytes in file image.
	e.addValue(8, dataSize)            // Number of bytes in memory image.
	e.addValue(8, alignment)

	// Output the text segment
	e.addBytes(bin...)
	// Output the data segment
	e.addBytes()

}

// add bytes elf
func (e *elf) addBytes(b ...byte) {
	e.elfBytes = append(e.elfBytes, b...)
}

// add value bytes elf
func (e *elf) addValue(size int, val uint64) {

	switch size {
	case 1:
		e.elfBytes = append(e.elfBytes, uint8le(uint8(val))...)
	case 2:
		e.elfBytes = append(e.elfBytes, uint16le(uint16(val))...)
	case 4:
		e.elfBytes = append(e.elfBytes, uint32le(uint32(val))...)
	case 8:
		e.elfBytes = append(e.elfBytes, uint64le(uint64(val))...)
	}

}

// save elf file
func (e *elf) saveBinFile() {
	err := os.WriteFile("./a.out", e.elfBytes, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("save...")
}
