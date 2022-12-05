package classfile

// Deprecated 和 Synthetic 结构体， 他们都是不包含任何数据，只有标记作用的结构体

type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

// 什么都不读取，空方法
func (self *MarkerAttribute) readInfo(reader *ClassReader) {}
