// Copyright (c) 2023 Parth Degama
// This code is licensed under MIT license

// register table

package x86_64

// register globle index
const (

	//
	noneReg = iota

	// 64 bit registers
	rax
	rcx
	rdx
	rbx
	rsp
	rbp
	rsi
	rdi
	r8
	r9
	r10
	r11
	r12
	r13
	r14
	r15

	// 32 bit registers
	eax
	ecx
	edx
	ebx
	esp
	ebp
	esi
	edi
	r8d
	r9d
	r10d
	r11d
	r12d
	r13d
	r14d
	r15d

	// 16 bit registers
	ax
	cx
	dx
	bx
	sp
	bp
	si
	di
	r8w
	r9w
	r10w
	r11w
	r12w
	r13w
	r14w
	r15w

	// 8 bit registers
	al
	cl
	dl
	bl
	ah
	ch
	dh
	bh
	spl
	bpl
	sil
	dil
	r8b
	r9b
	r10b
	r11b
	r12b
	r13b
	r14b
	r15b

	//
)

var registers []register = []register{

	// 64 bit registers
	{"RAX", rax, 64, 0, 0},
	{"RCX", rcx, 64, 1, 1},
	{"RDX", rdx, 64, 2, 2},
	{"RBX", rbx, 64, 3, 3},
	{"RSP", rsp, 64, 4, 4},
	{"RBP", rbp, 64, 5, 5},
	{"RSI", rsi, 64, 6, 6},
	{"RDI", rdi, 64, 7, 7},
	{"R8", r8, 64, 8, 0},
	{"R9", r9, 64, 9, 1},
	{"R10", r10, 64, 10, 2},
	{"R11", r11, 64, 11, 3},
	{"R12", r12, 64, 12, 4},
	{"R13", r13, 64, 13, 5},
	{"R14", r14, 64, 14, 6},
	{"R15", r15, 64, 15, 7},

	// 32 bit registers
	{"EAX", eax, 32, 0, 0},
	{"ECX", ecx, 32, 1, 1},
	{"EDX", edx, 32, 2, 2},
	{"EBX", ebx, 32, 3, 3},
	{"ESP", esp, 32, 4, 4},
	{"EBP", ebp, 32, 5, 5},
	{"ESI", esi, 32, 6, 6},
	{"EDI", edi, 32, 7, 7},
	{"R8D", r8d, 32, 8, 0},
	{"R9D", r9d, 32, 9, 1},
	{"R10D", r10d, 32, 10, 2},
	{"R11D", r11d, 32, 11, 3},
	{"R12D", r12d, 32, 12, 4},
	{"R13D", r13d, 32, 13, 5},
	{"R14D", r14d, 32, 14, 6},
	{"R15D", r15d, 32, 15, 7},

	// 16 bit registers
	{"AX", ax, 16, 0, 0},
	{"CX", cx, 16, 1, 1},
	{"DX", dx, 16, 2, 2},
	{"BX", bx, 16, 3, 3},
	{"SP", sp, 16, 4, 4},
	{"BP", bp, 16, 5, 5},
	{"SI", si, 16, 6, 6},
	{"DI", di, 16, 7, 7},
	{"R8W", r8w, 16, 8, 0},
	{"R9W", r9w, 16, 9, 1},
	{"R10W", r10w, 16, 10, 2},
	{"R11W", r11w, 16, 11, 3},
	{"R12W", r12w, 16, 12, 4},
	{"R13W", r13w, 16, 13, 5},
	{"R14W", r14w, 16, 14, 6},
	{"R15W", r15w, 16, 15, 7},

	// 8 bit registers
	{"AL", al, 8, 0, 0},
	{"CL", cl, 8, 1, 1},
	{"DL", dl, 8, 2, 2},
	{"BL", bl, 8, 3, 3},
	{"AH", ah, 8, 0, 4},
	{"CH", ch, 8, 1, 5},
	{"DH", dh, 8, 2, 6},
	{"BH", bh, 8, 3, 7},
	{"SPL", spl, 8, 4, 4},
	{"BPL", bpl, 8, 5, 5},
	{"SIL", sil, 8, 6, 6},
	{"DIL", dil, 8, 7, 7},
	{"R8B", r8b, 8, 8, 0},
	{"R9B", r9b, 8, 9, 1},
	{"R10B", r10b, 8, 10, 2},
	{"R11B", r11b, 8, 11, 3},
	{"R12B", r12b, 8, 12, 4},
	{"R13B", r13b, 8, 13, 5},
	{"R14B", r14b, 8, 14, 6},
	{"R15B", r15b, 8, 15, 7},

	//
}
