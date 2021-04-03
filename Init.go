package main

import (
	"fmt"
	"github.com/robfig/config"
)

var (
	QQEmailURL string
	ImagesURL  string
	Config     *config.Config

	// 正则表达式
	ReQQEmail = `(\d+)@qq.com`
	ReqImages = "https?://[^\"]+?(\\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))"
)

/**
 * @Author KyrisWang
 * @Description // 初始化url 初始化管道
 * @Date 4:06 下午 2021/4/2
 * @Param
 * @return
 */

func init() {
	var err error
	Config, err = config.ReadDefault("../reptile/conf/my.conf")
	if err != nil {
		panic(err)
	}

	QQEmailURL, err = Config.String("reptile", "QQEmailURL")
	if err != nil || QQEmailURL == "" {
		QQEmailURL = "https://tieba.baidu.com/p/6051076813?red_tag=1573533731"
	}

	ImagesURL, err = Config.String("reptile", "ImagesURL")
	if err != nil || ImagesURL == "" {
		ImagesURL = "https://www.bizhizu.cn/shouji/tag-可爱/1.html"
	}
}

// 封装一个打印err的方法
func ErrorF(err error, why string) bool {
	if err != nil {
		fmt.Println(why, err)
		return false
	}
	return true
}
