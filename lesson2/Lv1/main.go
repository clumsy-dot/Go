package main

import "fmt"

type Product struct { //定义Product结构体
	Name  string
	Price float64
	Stock int
}

func TotalValue(Price float64, number int) float64 { //计算价格并返回值
	value := Price * float64(number)
	return value
}

func IsInStock(stock int) bool { //判断数量是否够
	if stock > 0 {
		return true
	} else {
		return false
	}
}

func Info(p Product) string { //查询商品
	return ("商品名称:" + p.Name + ",价格:" + fmt.Sprintf("%.2f", p.Price) + ",库存:" + fmt.Sprintf("%d", p.Stock))
}

func (p *Product) Restock(amount int) { //计算进货后的数量
	p.Stock += amount
}

func (p *Product) Sell(amount int) (success bool, message string) { //判断库存是否足够
	if p.Stock >= amount {
		p.Stock -= amount
		return true, "出售成功"
	} else {
		return false, "库存不足，出售失败"
	}
}

func main() {
	p := Product{
		"Go编程书",
		59.9,
		100,
	}
	for {
		fmt.Println("请选择操作:1.查看商品信息 2.补货 3.出售商品 4.退出")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println(Info(p))
		case 2:
			fmt.Println("请输入补货数量:")
			var amount int
			fmt.Scanln(&amount)
			p.Restock(amount)
			fmt.Println("补货成功，当前库存:", p.Stock)
		case 3:
			fmt.Println("请输入出售数量:")
			var amount int
			fmt.Scanln(&amount)
			if amount > p.Stock {
				fmt.Println("数量不够")
			} else {
				success, message := p.Sell(amount)
				fmt.Println(success, message, "剩余库存:", p.Stock)
			}
		case 4:
			fmt.Println("感谢使用，再见")
		}
	}
}
