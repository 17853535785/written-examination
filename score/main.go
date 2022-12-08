package main

import "fmt"

// 一般求数组中的前k个数或者第k大第k小的数，优先想到堆排序算法
func main() {
	var nums = []int{60, 98, 59, 88, 98, 60, 90, 85, 83, 91}
	var k = 5
	v1(nums, k)
	v2(nums, k)
}

// v1 只考虑数组排序结果取前k个个数
// 示例输出：[88 90 91 98 98]
func v1(nums []int, k int) {
	numLength := len(nums)
	BuildMaxHeap(nums, numLength) // 此时堆顶元素为最大值
	// 求数组中的前k个值（成绩单中最高的k个分数）
	// 求前k个数，则执行k次交换
	for i := 0; i < k; i++ {
		nums[0], nums[numLength-1] = nums[numLength-1], nums[0]
		numLength--
		// 将当前范围内最大值（堆顶值）与最后的值交换后，调整堆
		ModifyHeap(nums, 0, numLength)
	}
	fmt.Println(nums[numLength:])
}

// v2 考虑存在分数相同的变量
// 示例输出：[85 88 90 91 98 98]
// 示例输出：[98 91 90 88 85]
func v2(nums []int, k int) {
	var ans = make([]int, 0)
	numLength := len(nums)
	BuildMaxHeap(nums, numLength) // 此时堆顶元素为最大值
	// 求数组中的前k个值（成绩单中最高的k个分数）
	for k > 0 {
		if len(ans) == 0 || nums[0] != ans[len(ans)-1] { // 如果当前最大元素与上一个最大元素相同，则分数相同算作一个值
			k--
			ans = append(ans, nums[0])
		}
		nums[0], nums[numLength-1] = nums[numLength-1], nums[0]
		numLength--
		// 将当前范围内最大值（堆顶值）与最后的值交换后，调整堆
		ModifyHeap(nums, 0, numLength)
	}
	fmt.Println(nums[numLength:])
	fmt.Println(ans)
}
