package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

// 压入一个栈帧
func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

// 弹出一个栈帧
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

// 返回当前栈帧
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) NewFrame(locals uint16, stack uint16) *Frame {
	return newFrame(t, uint(locals), uint(stack))
}
