package main

import (
	"fmt"
	"strconv"
	"sync"
)

/**
 * @Author KyrisWang
 * @Description //TODO 爬虫爬取图片下载 主要用到了多进程
 * @Date 3:52 下午 2021/4/2
 * @Param
 * @return
 */

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：多个协程向管道中添加图片链接
// 3.任务统计协程：检查任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	wg            sync.WaitGroup
	chanImageUrls chan string
	chanWebUrls   chan string
)

func main() {
	chanImageUrls = make(chan string, 100000)
	chanWebUrls = make(chan string, 35)

	// 启动协程 抓取每个页面的图片url 并且放在图片url管道中
	for i := 1; i < 36; i++ {
		wg.Add(1)
		go AddImgUrlToChan("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	wg.Add(1)
	go CheckOk()

	// 启动下载协程，从管道中读取数据并且进行下载消费。
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go DownloadImages()
	}

	wg.Wait()
}

func CheckOk() {
	var count int
	count = 0
	for {
		url := <-chanWebUrls
		fmt.Printf("当前url:%s 已完成爬取url任务", url)
		count++
		if count == 35 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}
