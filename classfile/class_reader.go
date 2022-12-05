package classfile

import "encoding/binary"

// 存储读取类文件的结构，是一个byte数组
type ClassReader struct {
	data []byte
}

/**
读取u1的数据
*/
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]       // 取出第一个字符
	self.data = self.data[1:] // 消费第一个字符
	return val
}

/**
读取u2长度的数据
*/
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data) // 取出16位数据
	self.data = self.data[2:]                 // 消费前两个byte数据
	return val
}

/**
读取u4长度的数据
*/
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data) // 取出32位数据
	self.data = self.data[4:]                 // 消费前四个byte数据
	return val
}

/**
读取64位数据
*/
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data) // 取出64位数据
	self.data = self.data[8:]                 // 消费前八个byte数据
	return val
}

/**
连续读n个连续的的uint16数据，n由第一个uint16数据定义
*/
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/**
读取指定数量的字节
*/
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]    // 取出前n个byte
	self.data = self.data[n:] // 消费前n个byte
	return bytes
}
