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
	var ans int
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return a[0] < b[0]
	})
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		temp := 1 // 代表最少有一台会同时在线
		end := nums[i][1]
		for j := i + 1; j < len(nums); j++ {
			if nums[j][0] < end {
				temp++
				end = min(nums[j][1], end)
			} else {
				break
			}
		}
		ans = max(ans, temp)
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
