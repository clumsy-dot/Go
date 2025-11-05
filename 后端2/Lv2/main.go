package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(filesname string, ch chan<- string) { //传入filesname和只读channel
	defer wg.Done() //延迟-1
	fmt.Println("正在下载" + filesname)
	time.Sleep(time.Second)  //等待
	ch <- filesname + "下载完成" //将文件名和下载完成传入通道
}
func main() {
	ch := make(chan string, 3) //定义通道容量为3
	files := []string{"file1.zip", "file2.pdf", "file3.mp4"}
	wg.Add(3)                 //计数3个goroutine
	for file := range files { //遍历files
		go f1(files[file], ch) //实现goroutine
	}
	wg.Wait()           //等待
	close(ch)           //关闭通道
	for a := range ch { //遍历ch并打印
		fmt.Println(a)
	}
	if len(ch) == 0 { //判断ch通道中是否用完
		fmt.Println("全部下载完成")
	}
}
