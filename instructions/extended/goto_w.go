package extended

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// wide index(16位)的goto
type GOTO_W struct {
	offset int
}

func (g *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	// 读4字节
	g.offset = int(reader.ReadInt32())
}

func (g *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.offset)
}
