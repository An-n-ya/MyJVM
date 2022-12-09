package comparisons

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 比较int，然后判断条件进行跳转
type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (ifcond *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 != v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 < v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 <= v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 > v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
func (ifcond *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 >= v1 {
		base.Branch(frame, ifcond.Offset)
	}
}
