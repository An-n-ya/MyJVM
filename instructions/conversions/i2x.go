package conversions

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点数转换成其他格式

type I2D struct{ base.NoOperandsInstruction }
type I2F struct{ base.NoOperandsInstruction }
type I2L struct{ base.NoOperandsInstruction }

func (i2d *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	val := float64(d)
	stack.PushDouble(val)
}
func (i2f *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	val := float32(d)
	stack.PushFloat(val)
}
func (i2l *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	val := int64(d)
	stack.PushLong(val)
}
