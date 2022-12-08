install:
	go build -o myJava main.go
test:install ch4
#	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.lang.String


ch3:
	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.src.testMain

ch4:
	./myJava test

