package stack

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 赋值栈顶的指令

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (dup *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.Top()
	stack.PushSlot(top)
}
