package extended

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type IFNULL struct{ base.BranchInstruction }
type IFNONNULL struct{ base.BranchInstruction }

func (i *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}
func (i *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, i.Offset)
	}
}
