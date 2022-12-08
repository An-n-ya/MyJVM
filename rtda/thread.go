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
