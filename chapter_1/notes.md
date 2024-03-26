## Go学习笔记---（一）

### 基本特性

#### 变量/常量

- 导入包的内置变量，字母大写开头

- 函数返回类型和变量放在参数列表后面，函数传参先变量后类型、多个相同类型变量可以把类型放最后

  > 可以简化C当中的函数复杂性
  >
  > 例如，`int (*fp)(int a, int b);` 声明了一个函数指针`fp`，这个指针指向的函数接受两个`int`类型的参数并返回一个`int`类型的结果。如果参数之一也是一个函数，声明会变得更加复杂，如 `int (*fp)(int (*ff)(int x, int y), int b)`。这种情况下，理解和编写类型声明变得非常困难。

- 在变量赋值方面，我觉得很有意思，Golang像是pyhton和C这两者（即弱类型与强类型语言）的结合体

  它可以不声明变量类型，像python或js一样 `a:=1` 或`var a=1`(初始化阶段)

  也可以提前声明类型，像C一样`var a int =10`

  ![image-20240324220054979](https://typora-zrx.oss-cn-beijing.aliyuncs.com/img/2024/03/24/20240324-220057.png)

  即（隐式声明/显示声明）

- 函数允许多值返回，接收的时候 := 可以自动识别并赋值，（放在入参列表后面）

  看上去挺有趣的，因为可以提前为返回值进行操作/(赋值)

- 有个很逆天的iota，专门使用在const初始化中，在其之后每多声明一个变量，就会自增一次，甚至不重置的话，会自动赋值