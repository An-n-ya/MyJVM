install:
	go build -o myJava main.go
test:instal test3 test4 test5
#	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.lang.String


test3:install
	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.src.testMain

test4:install
	./myJava test

test5:install
	javac ./java/src/GaussTest.java
	./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.src.GaussTest