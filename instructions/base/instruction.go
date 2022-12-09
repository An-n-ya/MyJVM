package base

import "MyJVM/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// 下面创建一些"基质"结构体，方便有类似行为的结构继承

// NoOperandsInstruction 没有操作数的指令结构(方便其他没有操作数的指令继承)
type NoOperandsInstruction struct{}

func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 没有操作数  什么也不做
}

// BranchInstruction 跳转指令结构
type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// 分支指令结构的操作数就是偏移地址
	bi.Offset = int(reader.ReadInt16())
}

// Index8Instruction 索引8位数据的指令结构（在加载和存储"局部变量表"的时候会用到这种结构）
// 会从操作数中读取一个 int8 的整数索引
type Index8Instruction struct {
	Index uint
}

func (i8i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

// Index16Instruction 索引16位数据的指令结构（在加载和存储"运行时常量池"的时候会用到这种结构）
// 会从操作数中读取一个 int16 的整数索引
type Index16Instruction struct {
	Index uint
}

func (i16i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i16i.Index = uint(reader.ReadUint16())
}
