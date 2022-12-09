package stack

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type SWAP struct{ base.NoOperandsInstruction } // 交换栈顶的两个元素的位置

func (s *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
