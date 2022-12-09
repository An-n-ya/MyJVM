package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type IAND struct{ base.NoOperandsInstruction } // jvm中没有布尔量，只会转成int和long进行布尔运算
type LAND struct{ base.NoOperandsInstruction }

func (and *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
func (and *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
