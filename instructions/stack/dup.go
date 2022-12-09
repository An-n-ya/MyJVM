package stack

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 赋值栈顶的指令

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

func (dup *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.Top()
	stack.PushSlot(top)
}

// Execute 赋值一个slot到第二个slot下
func (dup *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopSlot()
	v1 := stack.PopSlot()
	stack.PushSlot(v2)
	stack.PushSlot(v1)
	stack.PushSlot(v2)
}

// Execute 赋值一个slot到第三个slot下
func (dup *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v3 := stack.PopSlot()
	v2 := stack.PopSlot()
	v1 := stack.PopSlot()
	stack.PushSlot(v3)
	stack.PushSlot(v1)
	stack.PushSlot(v2)
	stack.PushSlot(v3)
}

func (dup *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopSlot()
	v1 := stack.PopSlot()
	stack.PushSlot(v1)
	stack.PushSlot(v2)
	stack.PushSlot(v1)
	stack.PushSlot(v2)
}

// Execute 复制头两个slot到第3个slot下
func (dup *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v3 := stack.PopSlot()
	v2 := stack.PopSlot()
	v1 := stack.PopSlot()

	stack.PushSlot(v2)
	stack.PushSlot(v3)
	stack.PushSlot(v1)
	stack.PushSlot(v2)
	stack.PushSlot(v3)
}

// Execute 复制头两个slot到第4个slot下
func (dup *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v4 := stack.PopSlot()
	v3 := stack.PopSlot()
	v2 := stack.PopSlot()
	v1 := stack.PopSlot()

	stack.PushSlot(v3)
	stack.PushSlot(v4)
	stack.PushSlot(v1)
	stack.PushSlot(v2)
	stack.PushSlot(v3)
	stack.PushSlot(v4)
}
