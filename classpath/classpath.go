package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseClasspath(cpOption)
	return cp
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	// 先在 boot中找， 再在extend中找，最后在用户classpath中找
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.userClasspath.String()
}

/**
解析 boot 和 extend 类路径
*/
func (self *ClassPath) parseBootAndExtClasspath(option string) {
	jreDir := getJreDir(option).(string)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(option string) interface{} {
	if option != "" && exists(option) {
		return option
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Cannot find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) parseClasspath(option string) {
	if option == "" {
		// 如果cpOption为空，就置为当前目录
		option = "."
	}
	self.userClasspath = newEntry(option)
}
