# Go语言核心编程

## 第一章 基础知识
### 1. Go代码的基础解读

    1.源程序以.go为结尾
    2.默认为UTF-8的编码方式
    3.标识符区分大小写
    4.语句结尾的分号可以省略
    5.函数以func开头，函数体开头的 "{" 必须与func同一行，而不能另起一行
    6.字符串变量用双引号 ""括起来
    7.调用包里的方法用 . 访问符，比如fmt.Printf
    8.main函数所在的包名必须是main
    
### 2. Go编译环境
    // 编译
    go build hello.go
    // 运行
    ./hello
    
### 3. Go 关键字（25个）
     break  default func    interface   select
     case   defer   go  map struct
     chan   else    goto    package switch
     const  fallthrough if range    type
     continue   for import return   var


名称|作用
---|---
package | 定义包名
import | 导入包名
const | 常量声明
var | 变量声明
func | 函数声明
defer | 延迟执行
go | 并发语法糖
return | 函数返回
struct | 定义结构类型
interface | 定义接口类型
map | 定义或声明map类型关键字
chan | 声明或创建通道类型关键字
     
### 4. 变量的声明方式

    1.显式的完整声明
    var varName dataType [ = value]
    eg. var a int = 1
    
    2.短类型声明
    varName := value
    注: “:=” 只能出现在函数内，（包括在方法内），Go编译器会自动进行类型判断
    可同时进行多个变量的声明并赋值： a,b := 1, "hello"
    