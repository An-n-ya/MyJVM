test:
	go build -o myJava main.go
#	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.lang.String
	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.src.testMain
