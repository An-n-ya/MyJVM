package extended

import (
	"MyJVM/instructions/base"
	"MyJVM/instructions/loads"
	"MyJVM/instructions/math"
)

// wide指令用于将其他指令扩展成支持16位局部变量表索引的指令

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (wide *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x16: // lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x17: // fload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x18: // dload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x19: // aload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x36: // istore
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x37: // lstore
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x38: // fstore
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x39: // dstore
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x3a: // astore
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		wide.modifiedInstruction = inst
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
	case 0xa9: // ret
		panic("ret current unsupported")

	}

}
