package x86_64

import (
	"fmt"
	"os"
)

func appTest1() {

	c := 0

	for _, opcode := range archOpcodeList {

		r := 0
		m := 0

		for _, oper := range opcode.operands {
			if isRegOperand(oper.t) {
				r++
			}

			if isMemOperand(oper.t) {
				m++
			}
		}

		if r != 0 && m == 0 {

			c++
			fmt.Println(opcode.name, r, m, opcode.operands, "modrm")

		}

	}

	fmt.Println(c)

	os.Exit(0)
}
