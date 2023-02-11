


# 进阶语法


context
sync 并发与 channel
反射与 unsafe
网络编程与 SQL 编程

AST 编程和模板编程
    - 规则引擎



- 类型系统

绝大多数 pl 的类型系统都是类似的，会有声明类型、实际类型之类的区分。

在 Go 反射里面，一个实例可以看出两部分：
    - 值
    - 实际类型



- reflect.Type 和 reflect.Value 

反射的相关 API 都在 reflect 包，最核心的两个：
    
    - reflect.Type：用于操作值，部分值是可以被反射修改的
    - reflect.Value：用于操作类信息，类信息只能读取

    reflect.Type 可以通过 reflect.Value 得到，但是反过来不行。


示意图：略



- reflect.Kind

reflect 包有一个很强的假设：`你知道你操作的是什么 Kind`

Kind：Kind 是一个枚举值，用来判断操作的对应类型；例如，是否是指针、是否是数组、是否是切片等

所以 reflect 的方法，如果你调用得不对，它直接 panic。


`在调用 API 之前一定要先读注释，确认什么情况下可以调用！`


例如在 reflect.Type 这里，这三个方法都有对应的 Kind 必须是什么，否则 panic。

    // 返回 struct 的字段数量
    NumField() int
    // ...
    NumIn() int
    NumOut() int
 


 
    
