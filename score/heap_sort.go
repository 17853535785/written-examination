package main

// BuildMaxHeap 构建大顶堆
func BuildMaxHeap(nums []int, heapSize int) {
	// 从最后一个父节点开始调整（完全二叉树中最后一个父节点为 n/2 -1）,依次倒数第2个父节点......
	for i := heapSize/2 - 1; i >= 0; i-- {
		ModifyHeap(nums, i, heapSize)
	}
}

// ModifyHeap 调整堆
func ModifyHeap(nums []int, index int, heapSize int) {
	father, leftSon, rightSon := index, 2*index+1, 2*index+2
	// 选择左右子节点中较大的和父节点交换
	if leftSon < heapSize && nums[leftSon] > nums[father] {
		father = leftSon
	}
	if rightSon < heapSize && nums[rightSon] > nums[father] {
		father = rightSon
	}
	// 如果出现了交换情况，则继续调整
	if father != index {
		nums[father], nums[index] = nums[index], nums[father]
		ModifyHeap(nums, father, heapSize)
	}
}
