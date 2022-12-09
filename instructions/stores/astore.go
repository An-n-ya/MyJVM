package stores

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 存储long类型的数据

type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

func (store *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, store.Index)
}

func (store *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}
func (store *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}
func (store *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}
func (store *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
