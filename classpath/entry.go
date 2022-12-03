package classpath

import (
	"os"
	"strings"
)

/**
-classpath/-cp用来指定目录，示例如下
	java -cp path/to/classes
	java -classpath path/to/lib1.jar
	java -cp path/to/lib2.zip
	java -cp path/to/lib2.zip:path/to/lib1.jar
	java -cp lib/*
*/

const pathListSeparator = string(os.PathListSeparator) // path的分隔符，macOS下是冒号 :

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 处理有冒号:，即多个目录的情况
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		// 处理有通配符的情况
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		// 处理 jar zip 后缀的classpath
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
