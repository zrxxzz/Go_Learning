## Go

> 在进一步了解Golang 语法、语言特性前，我们来思考一个问题，为什么云原生平台往往使用Go进行相应的开发？
>
> - Go语言天生支持并发，内部的协程、通道，使得并发程序变得非常简单直观
> - Go和C/CPP一样属于编译型语言，可以直接编译成机器代码（二进制代码），高性能
> - Go简单易读，像之前提到的短变量初始都非常易懂

### 基本特性

### 官网教程学习

> 上次因为不太确定能否OC，Golang学习拖了一周，复习同时学一些新知识

#### 循环

- 有个实现Math.Sqrt的手工办法

  ```go
  func Sqrt(x float64) float64 {
  	z:=1.0;
  	for i:=1;i<10;i++{
  		z -= (z*z-x)/(2*z)
  		fmt.Println(z)
  	}
  	return z;
  }
  ```

- Golang的switch case 默认在每个case后面都加了break 因此不需要再break了

- 无条件的switch case 可以看作是if-else 集合

- defer 语句会将函数推迟到外层函数返回之后执行

  这样可以使用在资源回收的场景中，确保作用域结束后，资源得到及时释放

  每次使用完一个资源时就应该及时使用`defer`释放

#### 方法

>其实方法就是一类特殊的函数

可以有两种实现方式，正常的函数形式，或者带接收参数的方法形式

- `func Abs(a Test) return_type`---调用方式`Abs(a)`
- `func (a Test) Abs return_type`---调用方式`a.Abs()`

可以为非结构体（自定义一个类型）定义方法--- 比如给自己的int 定义一个新方法

初始化时，可以像使用类型转换一样`mytype(value)`

> 和CPP一样，方法和函数传参时，如果想改变本身的值，就应该传入指针
>
> 还可以像引用一样高效

当出现指针和方法时，最好将指针作为接收者，这样既能及时修改，又能在调用时兼容值的调用

#### 接口

还是和上次说的一样，属于抽象基类

> nil 空指针也可以调用方法
>
> ![image-20240407144139421](https://typora-zrx.oss-cn-beijing.aliyuncs.com/img/2024/04/07/20240407-144143.png)

- 空接口可以保存任意类型




#### Tips

- 在输出时，其实用Println 和Print 都可以直接实现格式化输出的结果，即Printf中的`%v`效果

### 仓库

[Github 仓库](https://github.com/zrxxzz/Go_Learning)