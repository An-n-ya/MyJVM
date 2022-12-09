package rtda

import "math"

type OperandStack struct {
	size  uint // 当前位置
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

func (o *OperandStack) PushFloat(val float32) {
	o.slots[o.size].num = int32(math.Float32bits(val))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	return math.Float32frombits(uint32(o.slots[o.size].num))
}

func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].num = int32(val)
	o.slots[o.size+1].num = int32(val >> 32)
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (o *OperandStack) PushDouble(val float64) {
	o.PushLong(int64(math.Float64bits(val)))
}

func (o *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(o.PopLong()))
}

func (o *OperandStack) PushRef(ref *Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref
	o.slots[o.size].ref = nil
	return ref
}

// PushSlot 不关心变量类型的栈指令
func (o *OperandStack) PushSlot(slot Slot) {
	o.slots[o.size] = slot
	o.size++
}

func (o *OperandStack) PopSlot() Slot {
	o.size--
	return o.slots[o.size]
}

// 返回栈顶元素 不弹出
func (o *OperandStack) Top() Slot {
	return o.slots[o.size-1]
}
