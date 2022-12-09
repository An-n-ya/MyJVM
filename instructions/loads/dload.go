package loads

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type DLOAD struct{ base.Index8Instruction }      // 将8位操作数(索引)压入操作数栈
type DLOAD0 struct{ base.NoOperandsInstruction } // 把常量表索引为1的变量压入操作数栈
type DLOAD1 struct{ base.NoOperandsInstruction }
type DLOAD2 struct{ base.NoOperandsInstruction }
type DLOAD3 struct{ base.NoOperandsInstruction }

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (dload *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, uint(dload.Index))
}

func (dload *DLOAD0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}
func (dload *DLOAD1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}
func (dload *DLOAD2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}
func (dload *DLOAD3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}
