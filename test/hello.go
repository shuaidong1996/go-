package main

import (
	"BinaryTree"
	"math/rand"
	"fmt"
)
/*自定义类型*/
type myNode struct {
	val int
}
func (node *myNode) IsEqual(data BinaryTree.Data) int{
	node2 := data.(*myNode)
	if node.val > node2.val{
		return 1
	}else if node.val < node2.val{
		return -1
	}
	return 0
}

func main(){
	tree := BinaryTree.NewTree()
	for i:=0;i<10;i++{
		val := rand.Intn(100)
		node := myNode{val}
		var intNode BinaryTree.Data = &node
		if tree.AddNode(intNode) == false{
			fmt.Println("插入",val,"失败")
		}
		fmt.Print(val,"  ")
	}
	fmt.Println("")
	tree.Show()

	//查找
	if tree.FindNode(BinaryTree.Data(&myNode{25})) == true{
		fmt.Println("找到了25")
	}else{
		fmt.Println("没找到了25")
	}
	if tree.FindNode(BinaryTree.Data(&myNode{50})) == true{
		fmt.Println("找到了50")
	}else{
		fmt.Println("没找到了50")
	}

	//修改
	if tree.ModNode(BinaryTree.Data(&myNode{25}),BinaryTree.Data(&myNode{50})){
		fmt.Println("成功把25修改为50")
	}
	if tree.FindNode(BinaryTree.Data(&myNode{50})) == true{
		fmt.Println("找到了50")
	}

	//获取切片
	arr := tree.GetArrNode()
	for _,val := range arr{
		fmt.Print(val,"  ")
	}
	fmt.Println("\n")

	//删除
	for _,val := range arr{
		if tree.DelNode(val) == true{
			fmt.Println("成功删除",val.(*myNode).val)
		}else{
			fmt.Println("删除",val.(*myNode).val,"失败")
		}
		tree.Show()
	}
}