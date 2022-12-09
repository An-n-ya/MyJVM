package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (rem *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	stack.PushInt(-v) // 把结果压入栈中
}

func (rem *DNEG) Execute(frame *rtda.Frame) {
	// 浮点数不用检查除数，因为浮点数有NaN和inf的表示
	stack := frame.OperandStack()
	v := stack.PopDouble()
	stack.PushDouble(-v)
}

func (rem *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopFloat()
	stack.PushFloat(-v)
}

func (rem *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(-v)
}
