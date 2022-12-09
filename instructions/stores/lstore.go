package stores

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 存储long类型的数据

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (lstore *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, lstore.Index)
}

func (lstore *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (lstore *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (lstore *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (lstore *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
