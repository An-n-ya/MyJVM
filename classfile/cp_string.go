package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	// 从常量池去查找string
	return self.cp.getUtf8(self.stringIndex)
}
