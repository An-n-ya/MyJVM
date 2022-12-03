MY DIY JVM FOLLOWING "自己动手写Java虚拟机（张秀宏）"

## Installation
```shell
$ go build -o myJava main.go
```

## Usage
### 编译
```shell
./myJava -help
```
output: `Usage: ./myJava [-option] class [args...]`

### 测试cmd
```shell
./myJava -cp foo/bar MyApp.class arg1 arg2
```
output: `classpath:foo/bar class:MyApp.class args:[arg1 arg2]`

### 测试classpath相关代码
```shell
$ ./myJava -Xjre "/Users/zhanghuanyu/Library/Java/JavaVirtualMachines/azul-1.8.0_352/Contents/Home/jre" java.lang.Object
```
output: `classpath:/Users/zhanghuanyu/RustProjects/MyJVM class:java.lang.Object args:[]
class data: [202 254 186 190 0 0 0 52 0 78 7 0 49 10 0 1 0 50 10 0 17 0 51 10 0...`

## 注意
jdk11之后就没有单独的jre了，用户可以使用`jlink`生成自己的自定义的java运行时，比如：如果只需要java.base的话
```shell
 jlink \
         --add-modules java.base \
         --strip-debug \
         --no-man-pages \
         --no-header-files \
         --compress=2 \
         --output javaruntime
```
然后可以用`javaruntime/bin/java`来运行代码了。但javaruntime里其实还是没有jre文件夹的，因为jdk11后，jre已经融入到二进制文件里了。
详细信息可以参考 https://stackoverflow.com/questions/53733312/where-is-jre-11
如果需要用到`jre`里的文件的话，最好使用jdk1.8。


#### 关于jdk的安装位置
在macOS下，一般是安装在`/Library/Java/JavaVirtualMachines`下，当然也有可能在用户的资源库下：`/Users/xxx/Library/Java/JavaVirtualMachines`