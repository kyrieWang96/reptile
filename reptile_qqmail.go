package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/**
 * @Author KyrisWang
 * @Description //爬取百度贴吧的扣扣邮箱 不作处理打印出来
 * @Date 2:36 下午 2021/4/2
 * @Param
 * @return
 */

func GetQQMail() {
	resp, err := http.Get(QQEmailURL) // 创建连接
	ErrorF(err, "Create http:")

	data, err := ioutil.ReadAll(resp.Body) // 获取网页所有的值
	ErrorF(err, "Read http Body:")

	str := string(data)

	re := regexp.MustCompile(ReQQEmail)
	result := re.FindAllStringSubmatch(str, -1)

	for _, qqemail := range result {
		fmt.Println("Email:", qqemail[0])
		fmt.Println("qq", qqemail[1])
	}
}
