package constants

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type BIPUSH struct{ val int8 }  // bipush从操作数中获取一个byte型整数，扩展成int推入操作数栈顶
type SIPUSH struct{ val int16 } // sipush从操作数中获取一个short型整数，扩展成int推入操作数栈顶

func (bipush *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	bipush.val = reader.ReadInt8()
}

func (bipush *BIPUSH) Execute(frame *rtda.Frame) {
	val := int32(bipush.val)
	frame.OperandStack().PushInt(val)
}

func (sipush *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	sipush.val = reader.ReadInt16()
}

func (sipush *SIPUSH) Execute(frame *rtda.Frame) {
	val := int32(sipush.val)
	frame.OperandStack().PushInt(val)
}
