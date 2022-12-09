package control

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32   // 默认跳转offset
	npairs        int32   // 有多少分支
	matchOffsets  []int32 // 每个分支对应的条件和跳转地址
}

func (ls *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding() // 跳过0~3字节的padding
	ls.defaultOffset = reader.ReadInt32()
	ls.npairs = reader.ReadInt32()
	ls.matchOffsets = reader.ReadInt32s(ls.npairs * 2)
}

func (ls *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt() // switch操作数
	for i := int32(0); i < ls.npairs*2; i += 2 {
		// 遍历每个索引  i每次加2
		if ls.matchOffsets[i] == key {
			// 如果和switch条件匹配上了, 进行跳转
			offset := ls.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	// 如果都没有匹配到，跳转到默认地址
	base.Branch(frame, int(ls.defaultOffset))
}
