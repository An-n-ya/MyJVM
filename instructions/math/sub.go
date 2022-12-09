package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type DSUB struct{ base.NoOperandsInstruction }
type FSUB struct{ base.NoOperandsInstruction }
type ISUB struct{ base.NoOperandsInstruction }
type LSUB struct{ base.NoOperandsInstruction }

func (rem *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result) // 把结果压入栈中
}

func (rem *DSUB) Execute(frame *rtda.Frame) {
	// 浮点数不用检查除数，因为浮点数有NaN和inf的表示
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

func (rem *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

func (rem *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
