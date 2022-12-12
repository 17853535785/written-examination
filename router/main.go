package main

import (
	"fmt"
	"sort"
)

// 按照示例：
// 输出： 这个时间有最多 4 台设备连接了路由器
func main() {
	// 实际上为左闭右开区间
	nums := [][2]int{{1, 9}, {2, 4}, {5, 6}, {3, 7}, {1, 8}}
	count := calNum(nums)
	fmt.Printf("这个时间有最多 %d 台设备连接了路由器", count)
}

func calNum(nums [][2]int) int {
	var sortedNum []int
	mark := make(map[int]int) // 用来标记区间的start点和end点
	for _, item := range nums {
		sortedNum = append(sortedNum, []int{item[0], item[1]}...) // 将目标数组转换为一维数组后面遍历使用
		mark[item[0]] = 0                                         // 代表是区间start
		mark[item[1]] = 1                                         // 代表是区间end
	}
	sort.Ints(sortedNum)
	fmt.Println(sortedNum)
	var ans int
	var temp int
	for _, num := range sortedNum {
		if mark[num] == 0 { // 如果是区间start的话，代表有路由器接入
			temp++
		} else { // 如果是区间end点的话，代表有路由器断开连接
			temp--
		}
		if ans < temp { // 每次遍历完后，保存当前的最大结果值
			ans = temp
		}
	}
	return ans
}
