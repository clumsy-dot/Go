package main

import (
	"fmt"
)

func main() {
	var a, b int
	var x, input string
	fmt.Println("欢迎使用Go语言计算器!")
	fmt.Println("请输入两个整数和一个操作符，进行四则运算。")
	fmt.Println("输入exit退出程序")
	for {
		fmt.Println("请输入第一个整数:")
		fmt.Scanln(&a)
		fmt.Println("请输入操作符:")
		fmt.Scanln(&x)
		fmt.Println("请输入第二个整数:")
		fmt.Scanln(&b)
		if b == 0 || x != "+" && x != "-" && x != "/" && x != "*" {
			fmt.Println("输入有误，请重新输入")
			continue
		}
		switch x {
		case "+":
			fmt.Printf("%d+%d=%d", a, b, a+b)
		case "-":
			fmt.Printf("%d-%d=%d", a, b, a-b)
		case "/":
			fmt.Printf("%d/%d=%d", a, b, a/b)
		case "*":
			fmt.Printf("%dx%d=%d", a, b, a*b)
		}
		fmt.Println("是否继续?(exit退出)")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("感谢使用，再见！！！")
			break
		} else {
			continue
		}
	}
}
