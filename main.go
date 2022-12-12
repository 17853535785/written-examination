package main

import "fmt"

func main() {
	nums := []int{4, 5, 3}
	ans := abc(nums)
	fmt.Println(ans)
}
func abc(nums []int) [][]int {
	var ans [][]int
	mark := make(map[int]bool) // 记录访问过
	if len(nums) == 0 {
		return ans
	}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		for i := 0; i < len(nums); i++ {
			if mark[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			mark[nums[i]] = true
			dfs(path)
			path = path[:len(path)-1]
			mark[nums[i]] = false
		}
	}
	dfs([]int{})
	return ans
}
