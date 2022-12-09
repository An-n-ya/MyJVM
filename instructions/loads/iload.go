package loads

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type ILOAD struct{ base.Index8Instruction }      // 将8位操作数(索引)压入操作数栈
type ILOAD0 struct{ base.NoOperandsInstruction } // 把常量表索引为1的变量压入操作数栈
type ILOAD1 struct{ base.NoOperandsInstruction }
type ILOAD2 struct{ base.NoOperandsInstruction }
type ILOAD3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (iload *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(iload.Index))
}

func (iload *ILOAD0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
func (iload *ILOAD1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
func (iload *ILOAD2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}
func (iload *ILOAD3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
