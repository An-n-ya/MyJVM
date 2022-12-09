package math

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

type ISHL struct{ base.NoOperandsInstruction }  // int 左位移
type ISHR struct{ base.NoOperandsInstruction }  // int 算数右位移
type IUSHR struct{ base.NoOperandsInstruction } // int 逻辑右位移
type LSHL struct{ base.NoOperandsInstruction }
type LSHR struct{ base.NoOperandsInstruction }
type LUSHR struct{ base.NoOperandsInstruction }

func (shl *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 取出两个操作数
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f // 位移量：只取低5位，因为32位数字最多向左移动32位
	result := v1 << s
	stack.PushInt(result)
}

func (shl *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 取出两个操作数
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f // 位移量：只取低6位，因为64位数字最多左移64位
	result := v1 << s
	stack.PushLong(result)
}

func (shl *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}
func (shl *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// 先转成无符号后右移，最后再转回有符号数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (shl *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}
func (shl *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	// 先转成无符号后右移，最后再转回有符号数
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
