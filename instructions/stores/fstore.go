package stores

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 存储long类型的数据

type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (store *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, store.Index)
}

func (store *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}
func (store *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}
func (store *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}
func (store *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
