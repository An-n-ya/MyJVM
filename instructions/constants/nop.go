package constants

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (nop *NOP) Execute(frame *rtda.Frame) {
	// 什么也不做
}
