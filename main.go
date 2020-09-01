package main

import (
	"biliblili_fans/models"
	wg "biliblili_fans/wg"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

/*
目的 爬取给定up主的所有粉丝数据，包含内容为
1. 粉丝id
2. 名称
3. 等级
4. 粉丝数
5. 获赞数
6. 播放数

思路
1. 实现一个方法，可以根据给定的id爬取所有的粉丝id
2. 实现一个方法，可以根据给定的id爬取所有的上述信息
*/

const (
	// 获取up粉丝id的链接 GET
	// 参数 vmid pn=页号 ps=页面大小 order=desc
	FANS_URL = "https://api.bilibili.com/x/relation/followers"

	// 获取up信息的链接
	// 参数 mid
	INFO_URL = "https://api.bilibili.com/x/space/acc/info"
)

var bwg wg.BiliWaitGroup

func main() {
	GetAllFansLevel(74429422)
	bwg.Wait()
}

func GetUserinfo(mid int64) (info *models.Info) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error")
		}
	}()
	request := gorequest.New()
	request.Get(INFO_URL).Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36").
		Param("mid", strconv.FormatInt(mid, 10)).
		EndStruct(&info)

	return
}

func GetFans(mid int64, page int) (fans models.Fans) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error")
		}
	}()
	request := gorequest.New()
	request.Get(FANS_URL).Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36").
		Param("vmid", strconv.FormatInt(mid, 10)).
		Param("pn", strconv.Itoa(page)).
		Param("ps", "20").
		Param("order", "desc").
		EndStruct(&fans)

	return
}

// 该函数用于获取所有的粉丝等级分布，储存字段（mid, name, level）
func GetAllFansLevel(mid int64) {
	// 计算页数
	total := GetFans(mid, 1).Data.Total
	major := total / 20
	min := total % 20
	if min > 0 {
		major++
	}

	for i := 1; i <= int(major); i++ {
		//爬取每一页粉丝列表
		fans := GetFans(mid, i)
		upList := fans.Data.List
		// 拿到每个粉丝id然后开个goroutine去爬 mid, name, level信息
		for _, each := range upList {
			mid := each.Mid // 注意这里必须要用中间变量存一下
			bwg.Run(func() {
				info := GetUserinfo(mid)
				// 对这个info 进行信息提取加保存
				fmt.Printf("mid: %d, name: %s, level: %d\n", info.Data.Mid, info.Data.Name, info.Data.Level)
			})
		}
	}

}
