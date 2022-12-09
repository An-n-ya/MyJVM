package stack

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type POP struct{ base.NoOperandsInstruction }  // 弹出一个单位Slot(四字节)的数据, 没有返回值
type POP2 struct{ base.NoOperandsInstruction } // 弹出两个单位的数据, 没有返回值

func (p *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
