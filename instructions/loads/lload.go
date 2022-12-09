package loads

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type LLOAD struct{ base.Index8Instruction }      // 将8位操作数(索引)压入操作数栈
type LLOAD0 struct{ base.NoOperandsInstruction } // 把常量表索引为1的变量压入操作数栈
type LLOAD1 struct{ base.NoOperandsInstruction }
type LLOAD2 struct{ base.NoOperandsInstruction }
type LLOAD3 struct{ base.NoOperandsInstruction }

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (iload *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(iload.Index))
}

func (iload *LLOAD0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}
func (iload *LLOAD1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}
func (iload *LLOAD2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}
func (iload *LLOAD3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
