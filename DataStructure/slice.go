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
	"runtime"
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

	// 申请内存空间,单位:字节,返回值:void *类型,不能参与运算,s.Data是内存空间的首地址值
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
	fmt.Println()
}

/*
	定义添加切片元素的方法
*/
func (s *Slice) Append(Data ...int) {
	// 容错校验
	if s == nil {
		return
	}
	// 判断是否需要扩容
	for len(Data)+s.Len > s.Cap {
		// 扩展容量为原来的2倍,存贮新的内存地址
		s.Data = C.realloc(s.Data, C.size_t(s.Cap)*2*8)
		// 更新容量
		s.Cap *= 2
	}
	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	// 偏移p到结尾处
	p += uintptr(s.Len) * 8
	// 循环将Data中的数据存入内存中
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}
	// 修改len
	s.Len += len(Data)
}

/*
	根据下标获取切片元素
*/
func (s *Slice) GetData(index int) int {
	// 容错校验
	if s == nil {
		return -1
	}
	if index < 0 || index >= s.Len {
		return -1
	}
	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	// 偏移p到index指定的元素位置
	p += uintptr(index) * TAG
	// 获取数据并返回
	return *(*int)(unsafe.Pointer(p))
}

/*
	已知元素,返回下标
*/
func (s *Slice) SearchData(data int) int {
	// 容错校验
	if s == nil {
		return -1
	}
	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	for i := 0; i < s.Len; i++ {
		if *(*int)(unsafe.Pointer(p)) == data {
			return i
		}
		p += TAG
	}
	return -1
}

/*
	根据下标删除切片元素
*/
func (s *Slice) Delete(index int) {
	// 容错校验
	if s == nil {
		return
	}
	if index < 0 || index >= s.Len {
		return
	}
	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	// 偏移p到index指定的元素位置
	p += uintptr(index) * TAG
	// 定义变量记录p指代的元素的下一个元素
	afterP := p
	// 循环从index到s.Len,依次完成后一个元素给前一个元素赋值
	for i := index; i < s.Len; i++ {
		afterP += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(afterP))
		p += TAG
	}
	// 修改s.Len,去除一个元素
	s.Len -= 1
}

/*
	根据index向切片中插入数据
*/
func (s *Slice) Insert(data int, index int) {
	// 容错校验
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.Len {
		return
	}
	// 如果插入的index在切片结尾
	if index == s.Len {
		s.Append(data)
		return
	}
	// 判断是否需要扩容
	for s.Len+1 > s.Cap {
		// 扩展容量为原来的2倍,存贮新的内存地址
		s.Data = C.realloc(s.Data, C.size_t(s.Cap)*2*8)
		// 更新容量
		s.Cap *= 2
	}

	// 将s.Data转换成可以计算的数值
	p := uintptr(s.Data)
	// 偏移p到index指定的元素位置
	p += uintptr(index) * TAG
	// 获取插入元素完成后的最后一个元素的位置
	temp := uintptr(s.Data) + uintptr(s.Len)*TAG
	// 循环将index之后的元素依次后移--前一个元素给后一个元素赋值
	for i := s.Len; i > index; i-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}
	// 循环结束后,将要插入的数据写入p对应的内存中
	*(*int)(unsafe.Pointer(p)) = data
	// 修改s.len
	s.Len++
}

// 定义销毁切片的方法
func (s *Slice) Destory() {
	// 容错校验
	if s == nil || s.Data == nil {
		return
	}
	C.free(s.Data)
	// 将当前内存空间置空,驱使GC工作
	s.Data = nil
	s.Len = 0
	s.Cap = 0
	s = nil

	// 手动调用GC
	runtime.GC()
}

func main() {
	slice := new(Slice)
	// 创建切片
	slice.Create(5, 5, 1, 2, 3, 4, 5)
	// 打印切片
	//slice.Print()

	// 添加元素
	slice.Append(6, 7, 8, 9, 10)
	slice.Print()
	fmt.Printf("长度:%d,容量:%d\n", slice.Len, slice.Cap)

	// 获取元素
	ret := slice.GetData(1)
	fmt.Println("ret=", ret)

	// 给定元素,获取下标值
	idx := slice.SearchData(8)
	fmt.Println("下标:", idx)

	//删除元素
	slice.Delete(5)
	slice.Print()
	fmt.Printf("长度:%d,容量:%d\n", slice.Len, slice.Cap)

	// 插入元素
	slice.Insert(666, 9)
	slice.Print()
	fmt.Printf("长度:%d,容量:%d\n", slice.Len, slice.Cap)

	// 销毁切片
	fmt.Println("销毁切片:")
	slice.Destory()
	slice.Print()
}
