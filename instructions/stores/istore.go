package stores

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 存储long类型的数据

type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (store *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, store.Index)
}

func (store *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}
func (store *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}
func (store *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}
func (store *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
