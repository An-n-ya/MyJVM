package comparisons

import (
	"MyJVM/instructions/base"
	"MyJVM/rtda"
)

// 浮点类型的比较，与整数不同，会多一中可能：NaN
type DCMPG struct{ base.NoOperandsInstruction } // 遇到NaN返回true
type DCMPL struct{ base.NoOperandsInstruction } // 遇到NaN返回false

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
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

func (cmp *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}
func (cmp *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
