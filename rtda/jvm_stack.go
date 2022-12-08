package rtda

type Stack struct {
	maxSize uint
	size    uint   // 当前有多少个栈帧
	_top    *Frame // 当前栈帧
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		// TODO: 栈溢出异常
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		// 将待压入的栈帧的lower指向当前栈帧（即frame的前一个栈帧)
		frame.lower = s._top
	}
	s._top = frame // 更新栈顶栈帧
	s.size++       // 栈帧数量加一
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		// 栈为空就不能继续pop了
		panic("jvm stack is empty!")
	}
	top := s._top
	s._top = top.lower // 置当前栈顶为上一个栈帧
	top.lower = nil    // 断开引用，方便GC
	s.size--           // 栈帧数量减一

	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}
