package loads

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type FLOAD struct{ base.Index8Instruction }      // 将8位操作数(索引)压入操作数栈
type FLOAD0 struct{ base.NoOperandsInstruction } // 把常量表索引为1的变量压入操作数栈
type FLOAD1 struct{ base.NoOperandsInstruction }
type FLOAD2 struct{ base.NoOperandsInstruction }
type FLOAD3 struct{ base.NoOperandsInstruction }

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (iload *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(iload.Index))
}

func (iload *FLOAD0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}
func (iload *FLOAD1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}
func (iload *FLOAD2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}
func (iload *FLOAD3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
