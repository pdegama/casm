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
	{"RAX", rax, 64, 0, 0, true},
	{"RCX", rcx, 64, 1, 1, true},
	{"RDX", rdx, 64, 2, 2, true},
	{"RBX", rbx, 64, 3, 3, true},
	{"RSP", rsp, 64, 4, 4, true},
	{"RBP", rbp, 64, 5, 5, true},
	{"RSI", rsi, 64, 6, 6, true},
	{"RDI", rdi, 64, 7, 7, true},
	{"R8", r8, 64, 8, 0, true},
	{"R9", r9, 64, 9, 1, true},
	{"R10", r10, 64, 10, 2, true},
	{"R11", r11, 64, 11, 3, true},
	{"R12", r12, 64, 12, 4, true},
	{"R13", r13, 64, 13, 5, true},
	{"R14", r14, 64, 14, 6, true},
	{"R15", r15, 64, 15, 7, true},

	// 32 bit registers
	{"EAX", eax, 32, 0, 0, false},
	{"ECX", ecx, 32, 1, 1, false},
	{"EDX", edx, 32, 2, 2, false},
	{"EBX", ebx, 32, 3, 3, false},
	{"ESP", esp, 32, 4, 4, false},
	{"EBP", ebp, 32, 5, 5, false},
	{"ESI", esi, 32, 6, 6, false},
	{"EDI", edi, 32, 7, 7, false},
	{"R8D", r8d, 32, 8, 0, true},
	{"R9D", r9d, 32, 9, 1, true},
	{"R10D", r10d, 32, 10, 2, true},
	{"R11D", r11d, 32, 11, 3, true},
	{"R12D", r12d, 32, 12, 4, true},
	{"R13D", r13d, 32, 13, 5, true},
	{"R14D", r14d, 32, 14, 6, true},
	{"R15D", r15d, 32, 15, 7, true},

	// 16 bit registers
	{"AX", ax, 16, 0, 0, false},
	{"CX", cx, 16, 1, 1, false},
	{"DX", dx, 16, 2, 2, false},
	{"BX", bx, 16, 3, 3, false},
	{"SP", sp, 16, 4, 4, false},
	{"BP", bp, 16, 5, 5, false},
	{"SI", si, 16, 6, 6, false},
	{"DI", di, 16, 7, 7, false},
	{"R8W", r8w, 16, 8, 0, true},
	{"R9W", r9w, 16, 9, 1, true},
	{"R10W", r10w, 16, 10, 2, true},
	{"R11W", r11w, 16, 11, 3, true},
	{"R12W", r12w, 16, 12, 4, true},
	{"R13W", r13w, 16, 13, 5, true},
	{"R14W", r14w, 16, 14, 6, true},
	{"R15W", r15w, 16, 15, 7, true},

	// 8 bit registers
	{"AL", al, 8, 0, 0, false},
	{"CL", cl, 8, 1, 1, false},
	{"DL", dl, 8, 2, 2, false},
	{"BL", bl, 8, 3, 3, false},
	{"AH", ah, 8, 0, 4, false},
	{"CH", ch, 8, 1, 5, false},
	{"DH", dh, 8, 2, 6, false},
	{"BH", bh, 8, 3, 7, false},
	{"SPL", spl, 8, 4, 4, true},
	{"BPL", bpl, 8, 5, 5, true},
	{"SIL", sil, 8, 6, 6, true},
	{"DIL", dil, 8, 7, 7, true},
	{"R8B", r8b, 8, 8, 0, true},
	{"R9B", r9b, 8, 9, 1, true},
	{"R10B", r10b, 8, 10, 2, true},
	{"R11B", r11b, 8, 11, 3, true},
	{"R12B", r12b, 8, 12, 4, true},
	{"R13B", r13b, 8, 13, 5, true},
	{"R14B", r14b, 8, 14, 6, true},
	{"R15B", r15b, 8, 15, 7, true},

	//
}
