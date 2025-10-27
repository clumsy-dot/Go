package main

import "fmt"

func romeToInt(a string) int { //写函数输入a值并返回int类型的值
	rome := make(map[byte]int) //定义一个map并对其罗马数字输入值
	rome['I'] = 1
	rome['V'] = 5
	rome['X'] = 10
	rome['L'] = 50
	rome['C'] = 100
	rome['D'] = 500
	rome['M'] = 1000
	sum := 0                      //定义结果
	for i := 0; i < len(a); i++ { //遍历所输入的罗马数
		if i+1 < len(a) && rome[a[i]] < rome[a[i+1]] { //判断i+1会不会超出a的长度以保证不会报错并判断大小
			sum -= rome[a[i]]
		} else {
			sum += rome[a[i]]
		}
	}
	return sum //返回sum值
}

func main() {
	fmt.Println(romeToInt("III"))
	fmt.Println(romeToInt("MCMXCIV"))
	fmt.Println(romeToInt("LVII"))
	fmt.Println(romeToInt("IX"))
}
