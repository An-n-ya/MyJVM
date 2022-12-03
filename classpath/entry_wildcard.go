package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
处理目标路径下所有jar文件
*/
func newWildcardEntry(path string) CompositeEntry {
	// 获得 * 之前的字符
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			// 通配符不会对递归的目录进行处理
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry

}
