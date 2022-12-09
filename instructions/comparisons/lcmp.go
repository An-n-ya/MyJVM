package comparisons

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// long类型的比较 lcmp

type LCMP struct{ base.NoOperandsInstruction }

func (lcmp *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		// jvm中用int作为布尔值
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
