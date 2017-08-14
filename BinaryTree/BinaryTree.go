package BinaryTree

import (
	"fmt"
)

type Data interface{
	IsEqual(Data) int	//相等返回0  大于返回1 小于返回-1
}

type Node struct {
	left  *Node
	right *Node
	val  Data
}

type Tree struct{
	root *Node
}

func NewTree() *Tree {
	g_dong.left = nil
	return &Tree{nil}
}

func showTree(node *Node){
	if node != nil{
		fmt.Print(node.val,"  ")
		showTree(node.left)
		showTree(node.right)
	}
}
/*先序遍历*/
func (tree *Tree) Show(){
	showTree(tree.root)
	fmt.Println("")
}

func getArrSys(node *Node,arr *[]Data){
	if node != nil{
		*arr = append(*arr, node.val)
		getArrSys(node.left,arr)
		getArrSys(node.right,arr)
	}
}
/*获取切片*/
func (tree *Tree) GetArrNode() (arr []Data){
	arr = make([]Data,0)
	getArrSys(tree.root,&arr)
	return
}

func addNodeSys(node *Node,val Data) (*Node,bool) {
	if node == nil{
		return &Node{nil, nil, val},true
	}
	value := val.IsEqual(node.val)
	res := false
	if value == 0{
		return node,false
	}else if value == 1{
		node.right,res = addNodeSys(node.right,val)
	}else if value == -1{
		node.left,res = addNodeSys(node.left,val)
	}
	return node,res
}

/*增*/
func (tree *Tree) AddNode(val Data) bool{
	if tree == nil {
		return false
	}
	if tree.root == nil{
		tree.root = &Node{nil, nil, val}
		return true
	}
	value := val.IsEqual(tree.root.val)
	res:=false
	if value == 0{
		return false
	}else if value == 1{
		tree.root.right,res = addNodeSys(tree.root.right,val)
	}else if value == -1{
		tree.root.left,res = addNodeSys(tree.root.left,val)
	}
	return res
}

/*删*/
func (tree *Tree) DelNode(val Data) bool{
	/*如果删除的是根*/
	if val.IsEqual(tree.root.val) == 0{
		if tree.root.right != nil{
			/*如果有右孩子*/
			left := tree.root.left
			tree.root = tree.root.right
			curNode := tree.root
			parrent := tree.root
			for curNode != nil{
				parrent = curNode
				curNode = curNode.left
			}
			parrent.left = left
		}else {
			/*没有右孩子*/
			tree.root = tree.root.left
		}
		return true
	}
	/*不是根，先找到这个节点*/
	curNode := tree.root
	parrent := tree.root
	isLeft := true
	for curNode != nil{
		res := val.IsEqual(curNode.val)
		if res == 0 {
			/*找到了*/
			if curNode.right != nil{
				/*如果有右孩子*/
				left := curNode.left		//保存左孩子
				curNode.left = nil
				curNode = curNode.right
				cur := curNode
				par := cur
				for cur != nil{
					par = cur
					cur = cur.left
				}
				par.left = left
			}else{
				/*没有右孩子*/
				curNode = curNode.left
			}

			if isLeft == true{
				/*在父亲的左边*/
				parrent.left = curNode
			}else{
				/*右边*/
				parrent.right = curNode
			}
			return true
		}else if res == 1{
			parrent = curNode
			curNode = curNode.right
			isLeft = false
		}else {
			parrent = curNode
			curNode = curNode.left
			isLeft = true
		}
	}
	return false
}

/*改*/
func (tree *Tree) ModNode(val Data,newVal Data) bool{
	node := tree.root
	for node != nil{
		res := val.IsEqual(node.val)
		if res == 0{
			node.val = newVal
			return true
		}else if res == 1{
			node = node.right
		}else {
			node = node.left
		}
	}
	return false
}

/*查*/
func (tree *Tree) FindNode(val Data) bool{
	node := tree.root
	for node != nil{
		res := val.IsEqual(node.val)
		if res == 0{
			return true
		}else if res == 1{
			node = node.right
		}else {
			node = node.left
		}
	}
	return false
}
