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

// 单向链表逆序
func (node *LinkNode) reverse01() *LinkNode {
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
func (node *LinkNode) reverse02() *LinkNode {
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
