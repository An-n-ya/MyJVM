package classfile

type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc          uint16
	length           uint16
	name_index       uint16
	descriptor_index uint16
	index            uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	self.LocalVariableTable = make([]*LocalVariableEntry, length)
	for i := range self.LocalVariableTable {
		self.LocalVariableTable[i] = &LocalVariableEntry{
			startPc:          reader.readUint16(),
			length:           reader.readUint16(),
			name_index:       reader.readUint16(),
			descriptor_index: reader.readUint16(),
			index:            reader.readUint16(),
		}
	}
}
