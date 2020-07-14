package BST

import (
	"bytes"
	"fmt"
	"learningNote/数据结构/Utils/Interfaces"
)

//二分搜索叔
type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

type BST struct {
	root *Node
	size int
}

func New() *BST{
	return &BST{
		root: nil,
		size: 0,
	}
}

func (b *BST) GetSize() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素 e
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

// 向以 Node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (b *BST) add(n *Node, e interface{}) *Node {
	if n == nil {
		b.size++
		return &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
	}

	if Interfaces.Compare(n.e, e) < 0 {
		n.right = b.add(n.right, e)
	} else if Interfaces.Compare(n.e, e) > 0 {
		n.left = b.add(n.left, e)
	}
	return n
}

func (b *BST) Contains(e interface{}) bool {
	return contains(b.root, e)
}
func contains(node *Node, e interface{}) bool {
	if node == nil {
		return false
	}
	if Interfaces.Compare(node.e, e) == 0 {
		return true
	} else if Interfaces.Compare(node.e, e) < 0 {
		return contains(node.right, e)
	} else {
		return contains(node.left, e)
	}
}

func (b *BST) PreOrder() {
	preOrder(b.root)
}
func preOrder(n *Node) {
	if n == nil {
		return
	}
	println(n.e)
	fmt.Println(n.left)
	fmt.Println(n.right)
}

func (b *BST) InOrder() {
	inOrder(b.root)
}
func inOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.left)
	println(n.e)
	fmt.Println(n.right)
}

func (b *BST) PostOrder() {
	postOrder(b.root)
}
func postOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.left)
	fmt.Println(n.right)
	println(n.e)
}

// 寻找二分搜索树的最小元素
func (b *BST) Minimum() interface{} {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return minimum(b.root).e
}

// 返回以 Node 为根的二分搜索树的最小值所在的节点
func minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

func (b *BST) RemoveMin() interface{} {
	if b.root == nil {
		panic("is no tree")
	}
	min := b.Minimum()
	b.root = b.removeMin(b.root)
	return min
}

func (b *BST) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		b.size--
		return rightNode
	}
	n.left = b.removeMin(n.left)
	return n
}

func (b *BST) Remove(e interface{}) {
	b.root = b.remove(b.root, e)
}

func (b *BST) remove(n *Node, e interface{}) *Node {
	if n == nil {
		return nil
	}

	if Interfaces.Compare(n.e, e) < 0 {
		n.right = b.remove(n.right, e)
		return n
	} else if Interfaces.Compare(n.e, e) > 0 {
		n.left = b.remove(n.left, e)
		return n
	} else {
		//如果右节点为空
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			b.size--
			return leftNode
		}
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			b.size--
			return rightNode
		}
		//左右节点都不为空，从右节点找一个最小的值
		minNode := minimum(n.right)
		minNode.left = n.left
		minNode.right = b.removeMin(n.right)
		n.left = nil
		n.right = nil
		return minNode
	}
}

func (b *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(b.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
