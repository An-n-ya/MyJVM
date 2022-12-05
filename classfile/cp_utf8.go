package classfile

// ConstantInfo 结构体

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	// java用的是MUTF8, 和正规的utf8有点不一样, 需要解析
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	// TODO: 先使用string之间转换，细节之后再处理
	return string(bytes)
}
