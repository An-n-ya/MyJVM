MY DIY JVM FOLLOWING "自己动手写Java虚拟机（张秀宏）"

## Installation
```shell
$ go build -o myJava main.go
```

## Usage
```shell
./myJava -help
```
output: `Usage: ./myJava [-option] class [args...]`


```shell
./myJava -cp foo/bar MyApp.class arg1 arg2
```
output: `classpath:foo/bar class:MyApp.class args:[arg1 arg2]`
