package control

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffset    []int32
}

func (t *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()                 // tableswitch后面有0~3字节的padding
	t.defaultOffset = reader.ReadInt32() // 默认跳转的offset
	t.low = reader.ReadInt32()           // 下界
	t.high = reader.ReadInt32()          // 上界
	jumpOffsetsCount := t.high - t.low + 1
	// 往后读jumpOffsetsCount个32位字节码
	t.jumpOffset = reader.ReadInt32s(jumpOffsetsCount)
}

func (t *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt() // 获得switch操作数

	var offset int
	if index >= t.low && index <= t.high {
		// 如果是在上下界之内，直接按索引跳转
		offset = int(t.jumpOffset[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}
	base.Branch(frame, offset)
}
