package comparisons

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点类型的比较，与整数不同，会多一中可能：NaN
type FCMPG struct{ base.NoOperandsInstruction } // 遇到NaN返回true
type FCMPL struct{ base.NoOperandsInstruction } // 遇到NaN返回false

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		// 如果不属于以上三种情况，说明操作数中出现了NaN
		// 这时候把gFlag压入栈中
		if gFlag {
			stack.PushInt(1)
		} else {
			stack.PushInt(-1)
		}
	}
}

func (cmp *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}
func (cmp *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
