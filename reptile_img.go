package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author KyrisWang
 * @Description //图片爬取方法
 * @Date 3:53 下午 2021/4/2
 * @Param
 * @return
 */

// 获取网页最大url

// 获取图片地址 正则匹配图片地址
func ReptileImages(url string) []string {
	req, err := http.Get(url)
	if !ErrorF(err, "访问url失败："+url) {
		return nil
	}
	defer req.Body.Close()

	datas, err := ioutil.ReadAll(req.Body)
	if !ErrorF(err, "读取网页内容失败：") {
		return nil
	}

	re := regexp.MustCompile(ReqImages)
	imgUrls := re.FindAllString(string(datas), -1)
	fmt.Printf("共找到%d条结果\n", len(imgUrls))

	return imgUrls
}

//根据传来的图片地址下载
func DownloadImages() {
	for imgUrl := range chanImageUrls {
		fileName := GetFilenameFromUrl(imgUrl)
		ok := DownloadImage(imgUrl, fileName)
		if ok {
			fmt.Printf("URL:%s,下载成功", imgUrl)
		} else {
			fmt.Printf("URL:%s,下载失败", imgUrl)
		}
	}
	wg.Done()
}

func DownloadImage(imgUrl string, fileName string) bool {
	resp, err := http.Get(imgUrl)
	if !ErrorF(err, "get img false:") {
		return false
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if !ErrorF(err, "read img false:") {
		return false
	}

	name := "/Users/wangzhang/go/爬虫图片下载/" + fileName

	err = ioutil.WriteFile(name, bytes, 0666)
	if !ErrorF(err, "save img false: ") {
		return false
	}
	return true
}

// 截取url名字,返回图片名称
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 传入页面 URL正则 解析出来的图片URL放到管道中去
func AddImgUrlToChan(url string) {
	imgUrls := ReptileImages(url)
	for _, imgUrl := range imgUrls {
		chanImageUrls <- imgUrl
	}

	// 检测该页面是否爬完，爬完之后web—url放在管道里面
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanWebUrls <- url
	//

	wg.Done()
}
