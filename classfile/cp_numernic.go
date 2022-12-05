package classfile

import "math"

// 数值类型, ConstantInfo的Implementation

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 读取一个uint32位的数据，然后转换成go的int类型
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// 读取一个uint32位的数据，然后转换成go的int类型
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	// 读取一个uint32位的数据，然后转换成go的int类型
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	// 读取一个uint32位的数据，然后转换成go的int类型
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
