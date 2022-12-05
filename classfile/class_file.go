package classfile

import "fmt"

// class文件的结构
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData} // 初始化Reader
	cf = &ClassFile{}
	cf.read(cr) // 启动解析ClassFile过程
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	// My Very Cute Animal Turns In Savage Full Moon Area
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// 检查魔数
		// TODO: 需要用Java的exception报错，暂时先用panic
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	// Java1.8的主版本号是 0x34=52
	// 只有1.x版本会有minorversion
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	// 常量池里有thisClass引用的原始值 即为类名（类限定名）
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 运行到这里的只会是java.lang.Object， 它没有父类
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, classpathIdx := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(classpathIdx)
	}
	return interfaceNames
}
