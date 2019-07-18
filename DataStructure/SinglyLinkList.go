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

func main() {
	list := new(LinkNode)
	// 创建链表
	list.Create(1, 2, 3, 4, 5)
	// 打印链表
	// 递归法
	list.Print1()
	fmt.Println()
	// 循环法
	list.Print2()
}
