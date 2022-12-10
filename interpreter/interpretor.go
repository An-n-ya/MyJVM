package interpreter

import (
	"MyJVM/classfile"
	"MyJVM/instructions/base"
	"MyJVM/rtda"
	"fmt"
)

func interpret(methodInfo *classfile.MemberInfo) {
	// 从memberInfo中获取code属性
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code() // 拿到字节码

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
	// TODO: 暂时没有return语句，用catchErr做代替
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 解码
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// 执行
		fmt.Printf("pc:%2d inst:%T %V\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
