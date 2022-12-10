package cmd

import (
	"MyJVM/classfile"
	"MyJVM/classpath"
	"MyJVM/interpreter"
	"MyJVM/rtda"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}

/**
处理命令行里的flag
*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func Start() {
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

func startJVM(cmd *Cmd) {
	// 第五章测试
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

	// // 第四章测试
	// frame := rtda.NewFrame(100, 100)
	// testLocalVars(frame.LocalVars())
	// testOperandStack(frame.OperandStack())

	// 第三章测试
	// cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	//
	// className := strings.Replace(cmd.class, ".", "/", -1) // 替换所有的.为/
	//
	// cf := loadClass(className, cp)
	//
	// fmt.Println(cmd.class)
	// printClassInfo(cf)

	// 第二章测试
	// // 查找class
	// classData, _, err := cp.ReadClass(className)
	// if err != nil {
	// 	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	// 	return
	// }
	//
	// fmt.Printf("class data: %v\n", classData)
}

func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	// 找main方法
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constant count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}
	fmt.Printf("method count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, -3.141592653589793)
	vars.SetRef(9, nil)
	fmt.Println(vars.GetInt(0))
	fmt.Println(vars.GetInt(1))
	fmt.Println(vars.GetLong(2))
	fmt.Println(vars.GetLong(4))
	fmt.Println(vars.GetFloat(6))
	fmt.Println(vars.GetDouble(7))
	fmt.Println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(-3.141592653589793)
	ops.PushRef(nil)
	fmt.Println(ops.PopRef())
	fmt.Println(ops.PopDouble())
	fmt.Println(ops.PopFloat())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopLong())
	fmt.Println(ops.PopInt())
	fmt.Println(ops.PopInt())
}
