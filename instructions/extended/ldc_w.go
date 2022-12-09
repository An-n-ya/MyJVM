package extended

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type LDC_W struct {
	index uint
}

func (l *LDC_W) FetchOperands(reader *base.BytecodeReader) {
	// 这里变成读取16位
	l.index = uint(reader.ReadUint16())
}

func (l *LDC_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// TODO: 检查这里是否正确，是getint么
	val := frame.LocalVars().GetInt(l.index)
	stack.PushInt(val)
}
