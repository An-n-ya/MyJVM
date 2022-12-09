package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }

func (rem *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		// 除数不能为零
		panic("java.lang.ArithmeticException: div by zero")
	}
	result := v1 / v2
	stack.PushInt(result) // 把结果压入栈中
}

func (rem *DDIV) Execute(frame *rtda.Frame) {
	// 浮点数不用检查除数，因为浮点数有NaN和inf的表示
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

func (rem *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

func (rem *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		// 除数不能为零
		panic("java.lang.ArithmeticException: div by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}
