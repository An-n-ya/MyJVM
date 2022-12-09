package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
	"math"
)

type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (rem *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		// 除数不能为零
		panic("java.lang.ArithmeticException: div by zero")
	}
	result := v1 % v2
	stack.PushInt(result) // 把结果压入栈中
}

func (rem *DREM) Execute(frame *rtda.Frame) {
	// 浮点数不用检查除数，因为浮点数有NaN和inf的表示
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

func (rem *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	tmp := math.Mod(float64(v1), float64(v2))
	result := float32(tmp)
	stack.PushFloat(result)
}

func (rem *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: div by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}
