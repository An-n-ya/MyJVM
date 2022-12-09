package conversions

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点数转换成其他格式

type L2D struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }
type L2F struct{ base.NoOperandsInstruction }

func (l2d *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	val := float64(d)
	stack.PushDouble(val)
}
func (l2i *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	val := int32(d)
	stack.PushInt(val)
}
func (l2l *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	val := float32(d)
	stack.PushFloat(val)
}
