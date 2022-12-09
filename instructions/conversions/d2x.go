package conversions

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点数转换成其他格式

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (d2i *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	val := float32(d)
	stack.PushFloat(val)
}
func (d2i *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	val := int32(d)
	stack.PushInt(val)
}
func (d2i *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	val := int64(d)
	stack.PushLong(val)
}
