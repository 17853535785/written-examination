package main

import "fmt"

type TreeNode struct {
	Val    string
	Father *TreeNode // 指向父节点
	Left   *TreeNode // 指向左孩子节点
	Right  *TreeNode // 指向右孩子节点
}

// 按照示例，输入root,A2,B1
// 输出：学生A2和学生B1的批准人为：院长1
func main() {
	root := TreeNode{}
	proA := TreeNode{}
	proB := TreeNode{}
	teaA := TreeNode{}
	teaB := TreeNode{}
	stuA1 := TreeNode{
		Val:    "学生A1",
		Father: &teaA,
		Left:   nil,
		Right:  nil,
	}
	stuA2 := TreeNode{
		Val:    "学生A2",
		Father: &teaA,
		Left:   nil,
		Right:  nil,
	}
	stuB1 := TreeNode{
		Val:    "学生B1",
		Father: &teaB,
		Left:   nil,
		Right:  nil,
	}
	stuB2 := TreeNode{
		Val:    "学生B2",
		Father: &teaB,
		Left:   nil,
		Right:  nil,
	}
	teaA.Val = "班主任A"
	teaA.Father = &proA
	teaA.Left = &stuA1
	teaA.Right = &stuA2
	teaB.Val = "班主任B"
	teaB.Father = &proA
	teaB.Left = &stuB1
	teaB.Right = &stuB2
	proA.Val = "院长1"
	proA.Father = &root
	proA.Left = &teaA
	proA.Right = &teaB
	proB.Val = "院长2"
	proA.Father = &root
	root.Val = "校长"
	root.Father = nil
	root.Left = &proA
	root.Right = &proB
	ans := searchTarget(&root, &stuA2, &stuB1)
	fmt.Printf("%s和%s的批准人为：%s", stuA2.Val, stuB1.Val, ans.Val)
}

// searchTarget 寻找两个节点的最近公共祖先
func searchTarget(root, nodeA, nodeB *TreeNode) *TreeNode {
	// 因为节点本身带有指向父节点的指针，所以可以直接向上遍历
	// 思路：开始从A节点不断往上遍历并记录已经访问过的祖先节点，
	mark := make(map[*TreeNode]struct{})
	for nodeA != nil {
		mark[nodeA] = struct{}{}
		nodeA = nodeA.Father
	}
	// 再从B节点不断往上遍历，遇到已经访问过的既是要寻找的
	for nodeB != nil {
		if _, ok := mark[nodeB]; ok {
			return nodeB
		}
		nodeB = nodeB.Father
	}
	return nil
}
