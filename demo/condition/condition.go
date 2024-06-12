package main // 声明main包，表示这是一个可执行程序

import (
	"fmt"     // 导入fmt包，用于格式化输出和读取输入
	"strings" // 导入strings包，用于字符串操作
)

func main() { // main函数，程序的入口点
	var s string // 声明一个字符串变量s，用于存储用户输入

	fmt.Scanln(&s) // 使用fmt.Scanln函数读取用户输入，直到遇到换行符
	// Scanln是Scan的变体，它在读取到换行符时会停止读取，并将换行符从输入中移除

	if strings.HasPrefix(s, "yes") { // 使用strings.HasPrefix检查s是否以"yes"开头
		fmt.Println("Cool.") // 如果是，打印"Cool."
	} else if strings.HasPrefix(s, "no") { // 如果不是，再检查s是否以"no"开头
		fmt.Println("Tell me why.") // 如果是，打印"Tell me why."
	} else { // 如果s既不以"yes"也不以"no"开头
		fmt.Println("WTF?") // 打印"WTF?"
	}
}
