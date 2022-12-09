package conversions

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点数转换成其他格式

type F2D struct{ base.NoOperandsInstruction }
type F2I struct{ base.NoOperandsInstruction }
type F2L struct{ base.NoOperandsInstruction }

func (f2d *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	val := float64(d)
	stack.PushDouble(val)
}
func (f2i *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	val := int32(d)
	stack.PushInt(val)
}
func (f2l *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	val := int64(d)
	stack.PushLong(val)
}
