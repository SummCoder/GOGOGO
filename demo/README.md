## demo

go基础知识

### 运行第一个go程序

```go
package main

import(
	"fmt"
)

func main() {
	fmt.Println("Hello, Go World!")
}
```

执行`go run main.go`

```text
Hello, Go World!
```

解释第一个程序：
- 包名，源码开头需要使用`package $name`指定包名
- 引用其他包，使用`import`，`fmt`为`format`的缩写
- 入口函数，main包的main函数是程序执行的入口

### 条件控制：if-else

```go
func main() {
    var s 
    stringfmt.Scanln(&s)
    if strings.HasPrefix(s, "yes") {
        fmt.Println("Cool.")
    } else if strings.HasPrefix(s, "no") {
        fmt.Println("Tell me why.")
    } else {
        fmt.Println("WTF?")
    }
}
```

- var用以定义变量
- &s取s的指针
- strings包提供一些字符串操作函数
- 其他包方法名总是首字母大写，例如这里的`Scanln`
    - 大写：导出符号，其他包可以引用
    - 小写：私有符号
- 条件语句可以无括号

### 条件控制：switch

```go
func main() {
	var num int
	_ , err := fmt.Scan(&num)
	if err != nil {
		fmt.Printf("Err input num: %v", err)
		return
	}
	switch num {
		case 0:
			fmt.Println("OK")
		case 1, 2, 3:
			fmt.Println("Oh, small integers")
		case 4:
		case 5:
			fmt.Println("It's either 4 or 5")
			fallthrough
		case 6:
			fmt.Println("Guess what you'll get.")
		default:
			fmt.Println("This is the default branch.")
	}
}
```

- 下划线表示忽略返回值
- := 用于定义新变量并赋值
- switch语法
    - case可以包含多个value
    - 需要主动fallthrough而C语言需要主动break

- 输入4会得到什么输出？无输出
- 输入5会得到什么输出？
```text
It's either 4 or 5
Guess what you'll get.
```

### 循环控制：for

- 只有for，没有while，也没有do...while
- 注意for语句里定义的变量作用域
- `i = 1024`不能用`:=`，因为`i`已经定义过了

```go
i := 10
for ; i > 0; i-- {
    fmt.Println("i = ", i)
}
```

### 数据结构：数组&切片

- 数组：内存中连续存储、固定长度的同类型数据
- 切片：数组中的某一段引用？
- 左闭右开语法糖
- append方法会生成一个新的切片
- `slice`的数据结构`reflect.SliceHeader`
```go
type SliceHeader struct {
	Data 	uintptr	// 切片指向的开始地址
	Len		int		// 切片的长度
	Cap		int 	// 切片存储容量>=Len
}
```


```go
func exampleArrayAndSlice() {
	var arr [5]int
	for i :=0; i < len(arr); i++ {
		arr[i] = i
	}

	// 切片左闭右开
	var slice1 []int = arr[2:5]
	// 等同于slice1 := arr[2:5]?
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}

	slice2 := append(slice1, 6)
	slice2[0] = 0
	arr[2] = 100
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice2[i])
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("part1[%v] = %v\n", i, slice1[i])
	}
}
```

结果：
```text
part1[0] = 2
part1[1] = 3
part1[2] = 4
part1[0] = 0
part1[1] = 3
part1[2] = 4
part1[3] = 6
part1[0] = 100
part1[1] = 3
part1[2] = 4
```

由上述结果能得出什么结论？
- 切片是对数组的引用，故而改变原数组也会改变切片
- slice2为何不受影响？使用append函数对切片进行扩展时，如果切片的容量不足以容纳新元素，Go语言会创建一个新的数组，并将原始数组的内容复制到新数组中，然后在新数组中添加新元素。
- 多维数组？`var arr [5][6]int`



