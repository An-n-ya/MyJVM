package rtda

type Slot struct {
	num int32   // 数据值（只是用int32存储，具体可能代表不同的类型）
	ref *Object // 用于GC的引用
}
