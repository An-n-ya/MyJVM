package comparisons

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 要两个引用弹出来, 比较两个引用是否相等进行跳转

type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

func (ifcond *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 == v2 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v1 != v2 {
		base.Branch(frame, ifcond.Offset)
	}
}
