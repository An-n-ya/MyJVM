package cosntants

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type LDC struct {
	index uint
}

func (l *LDC) FetchOperands(reader *base.BytecodeReader) {
	l.index = uint(reader.ReadUint8())
}

func (l *LDC) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// TODO: 检查这里是否正确，是getint么
	val := frame.LocalVars().GetInt(l.index)
	stack.PushInt(val)
}
