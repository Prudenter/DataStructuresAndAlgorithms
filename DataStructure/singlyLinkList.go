/**
* @Time : 19-7-19 上午12:14
* @Author : ASlowPerson
* @File : SinglyLinkList.go
* @Software: GoLand
 */
package main

import "fmt"

/*
	单向链表
*/

// 定义链表的结点
type LinkNode struct {
	Data interface{} // 数据域
	Next *LinkNode   // 指针域
}

// 创建链表
func (node *LinkNode) Create(data ...interface{}) {
	// 容错校验
	if node == nil || data == nil {
		return
	}
	// 创建头结点
	head := node
	// 循环遍历data,依次取出数据创建单向链表
	for _, v := range data {
		// 创建新结点,并且初始化
		newNode := new(LinkNode)
		newNode.Data = v
		newNode.Next = nil

		// 将当前结点的下一个结点赋值为新结点
		node.Next = newNode
		// 更新新节点为当前结点
		node = node.Next
	}
	// 将node赋值为头结点
	node = head
}

// 打印链表--递归法
func (node *LinkNode) Print1() {
	// 容错校验,也是递归的出口,尾结点.next==nil
	if node == nil {
		return
	}
	if node.Data != nil {
		fmt.Print(node.Data, " ")
	}
	// 递归调用本函数,打印整个链表
	node.Next.Print1()
}

// 打印链表--循环法
func (node *LinkNode) Print2() {
	// 容错校验
	if node == nil {
		return
	}
	for node.Next != nil {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
	}
}

// 获取链表的长度
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}
	// 定义一个计数器
	i := 0
	// 循环统计链表中数据结点的个数
	for node.Next != nil {
		//后移node结点
		node = node.Next
		i++
	}
	return i
}

// 插入链表结点--头插法
func (node *LinkNode) InsertByHead(data interface{}) {
	if node == nil || data == nil {
		return
	}
	// 创建新结点, 初始化
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	// 将新结点的下一个结点,赋值为 原链表的第一个结点(node.Next)
	newNode.Next = node.Next

	// 将头结点的下一个结点赋值为新结点
	node.Next = newNode
}

// 插入链表结点--头插法
func (node *LinkNode) InsertByTail(data interface{}) {
	if node == nil || data == nil {
		return
	}
	// 创建新结点, 初始化
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	// 找到原链表的最后一个结点
	for node.Next != nil {
		node = node.Next
	}
	// 将新结点设置为尾结点
	node.Next = newNode
}

// 插入链表结点--按位置插
func (node *LinkNode) InsertByIndex(data interface{}, index int) {
	if node == nil || data == nil {
		return
	}
	if index < 0 || index > node.Length() {
		return
	}
	// 如果插入位置在尾部
	if index == node.Length() {
		node.InsertByTail(data)
		return
	}

	// 创建新结点
	newNode := new(LinkNode)
	newNode.Data = data
	newNode.Next = nil

	// 找寻带插入位置
	for i := 0; i < index-1; i++ {
		node = node.Next
	}
	// 将新结点指向node的下一个结点
	newNode.Next = node.Next
	// 将node结点指向newnode
	node.Next = newNode
}

// 单向链表逆序
func (node *LinkNode) reverse1() *LinkNode {
	fmt.Println(111111111)
	// node=head  prevNode=p1  nextnode=p2
	// 容错校验
	if node == nil {
		return nil
	}
	// 记录当前结点
	currentNode := node
	// 记录当前结点的下一个结点
	var nextNode *LinkNode

	for node.Next != nil {
		nextNode = node.Next
		// 记录下一个结点的下一个结点
		node.Next = nextNode.Next
		// 逆序操作,将下一个结点指向当前结点
		nextNode.Next = currentNode
		// 后移当前结点到下一个结点,循环逆序链表
		currentNode = nextNode
	}
	return nextNode
}

// 单向链表逆序
func (node *LinkNode) reverse2() *LinkNode {
	fmt.Println(222222222)
	// 容错校验
	if node == nil {
		return nil
	}
	// 记录当前结点
	currentNode := node
	// 记录当前结点的下一个结点
	nextNode := node.Next

	for {
		// 记录下一个结点的下一个结点
		node.Next = nextNode.Next
		// 逆序操作,将下一个结点指向当前结点
		nextNode.Next = currentNode
		// 后移当前结点到下一个结点,循环逆序链表
		currentNode = nextNode
		if node.Next == nil {
			break
		}
		nextNode = node.Next
	}
	return nextNode
}

func main() {
	list := new(LinkNode)
	// 创建链表
	list.Create(1, 2, 3, 4, 5, 6)
	// 打印链表
	// 递归法
	list.Print1()
	fmt.Println()
	// 循环法
	list.Print2()
	fmt.Println()

	// 获取链表长度
	ret := list.Length()
	fmt.Println("长度:", ret)

	/*	// 插入链表数据
		// 头插法
		list.InsertByHead(666)
		list.Print1()
		fmt.Println()

		// 尾插法
		list.InsertByTail(777)
		list.Print1()
		fmt.Println()

		// 按位置插入
		list.InsertByIndex(888, 8)
		list.Print1()
		fmt.Println()*/

	//re := list.reverse1()
	re := list.reverse2()
	re.Print1()
}
