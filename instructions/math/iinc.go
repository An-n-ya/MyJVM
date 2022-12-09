package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// iinc给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint  // 局部变量表中的索引
	Const int32 // 增加的数值
}

func (inc *IINC) FetchOperands(reader *base.BytecodeReader) {
	inc.Index = uint(reader.ReadUint8())
	inc.Const = int32(reader.ReadUint8())
}

func (inc *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(inc.Index)
	val += inc.Const
	localVars.SetInt(inc.Index, val)
}
