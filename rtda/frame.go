package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// 根据最大局部变量表 和 最大操作数栈 的数量创建Frame, 这两个关键信息由编译器得出
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (f *Frame) SetNextPC(pc int) {
	f.nextPC = pc
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}
