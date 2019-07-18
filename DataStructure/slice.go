/**
* @Time : 19-7-18 下午2:18
* @Author : ASlowPerson
* @File : slice.go
* @Software: GoLand
 */

package main

/*
#include <stdlib.h>
*/
import "C" // 添加头文件,导入C代码
import (
	"fmt"
	"unsafe"
)

// 定义切片类型
type Slice struct {
	Data unsafe.Pointer // go语言中的万能指针类型,没有具体数据类型,不能进行运算.
	Len  int            // 切片的数据元素个数
	Cap  int            // 可扩展的有效容量
}

const TAG = 8

/*
	定义创建切片的方法
*/
func (s *Slice) Create(l int, c int, Data ...int) {
	// 容错校验
	if s == nil || Data == nil {
		return
	}
	if len(Data) == 0 {
		return
	}
	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}

	// 申请内存空间,单位:字节,返回值:void *类型,不能参与运算,s.Data是一个地址值
	s.Data = C.malloc(C.size_t(c) * 8)
	// 初始化长度和容量
	s.Len = l
	s.Cap = c
	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	// 遍历Data集合,将数据逐个存入申请的内存中
	for _, v := range Data {
		// 将数值p转换回地址值,并转换为具体数据类型,解引用,赋值
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}
}

/*
	定义打印切片方法
*/
func (s *Slice) Print() {
	if s == nil {
		return
	}
	// 将地址转换为可以计算的数值
	p := uintptr(s.Data)
	// 按len循环打印切片元素
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += TAG
	}
}

func main() {
	slice := new(Slice)
	slice.Create(5, 5, 1, 2, 3, 4, 5)
	slice.Print()
}
