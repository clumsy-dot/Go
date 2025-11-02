package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(filesname string, ch chan<- string) {
	defer wg.Done()
	fmt.Println("正在下载" + filesname)
	time.Sleep(time.Second)
	ch <- filesname + "下载完成"
}
func main() {
	ch := make(chan string, 3)
	files := []string{"file1.zip", "file2.pdf", "file3.mp4"}
	wg.Add(3)
	for file := range files {
		go f1(files[file], ch)
	}
	wg.Wait()
	close(ch)
	for a := range ch {
		fmt.Println(a)
	}
	if len(ch) == 0 {
		fmt.Println("全部下载完成")
	}
}
