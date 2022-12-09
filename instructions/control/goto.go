package control

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (g *GOTO) Execute(frame *rtda.Frame) {
	// 直接去offset即可
	base.Branch(frame, g.Offset)
}
