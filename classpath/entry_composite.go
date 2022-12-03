package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历数组里的每个entry，每个entry都去寻找className
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	// 输出每个entry的字符串，然后用分隔符连接起来
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	// 按分隔符分隔后，挨个处理每个entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
