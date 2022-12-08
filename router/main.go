package main

import "fmt"

func main() {
	// 实际上为左闭右开区间
	nums := [][2]int{{1, 9}, {2, 4}, {5, 6}, {3, 7}, {1, 8}}
	var t = 5
	count := calNum(nums, t)
	fmt.Printf("在 %d 这个时间有 %d 台设备连接了路由器", t, count)
}

func calNum(nums [][2]int, t int) int {
	var ans int
	for _, period := range nums {
		if t >= period[0] && t < period[1] {
			ans++
		}
	}
	return ans
}
