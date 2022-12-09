package stores

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 存储long类型的数据

type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (store *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, store.Index)
}

func (store *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}
func (store *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}
func (store *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}
func (store *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
