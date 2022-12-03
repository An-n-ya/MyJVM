package cmd

import (
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("DIY JVM version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		// 如果是help或者没有传任何参数,就打印帮助信息
		printUsage()
	} else {
		startJVM(cmd)
	}
}
