package loads

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type ALOAD struct{ base.Index8Instruction }      // 将8位操作数(索引)压入操作数栈
type ALOAD0 struct{ base.NoOperandsInstruction } // 把常量表索引为1的变量压入操作数栈
type ALOAD1 struct{ base.NoOperandsInstruction }
type ALOAD2 struct{ base.NoOperandsInstruction }
type ALOAD3 struct{ base.NoOperandsInstruction }

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (load *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(load.Index))
}

func (load *ALOAD0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}
func (load *ALOAD1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}
func (load *ALOAD2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}
func (load *ALOAD3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
