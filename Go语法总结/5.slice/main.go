package main

import (
	"fmt"
)

func changeSliceDataCapError(data []int) {
	fmt.Println("函数中.............")
	println("第十个元素的地址", &data[9])
	println("函数中data的地址", &data)
	data[0] = 10
	data = append(data, 10)
	fmt.Println("执行后.............")
	println("函数中data的地址", &data)
	fmt.Println("data:", data)
	println("Len:", len(data))
	//cap不够的话会开辟一个新空间，变成两倍，把原来的复制过来，然后执行append
	//所以main函数中的data指向的slice对象 len没变，cap没变，只有data[0]变了
	println("Cap:", cap(data))
	println("第十个元素的地址", &data[9])
	println("第十一个元素的地址", &data[10])
}

//传递指针
func changeDataByPoint(data *[]int) {
	*data = append(*data, 10)
}

func main() {
	data := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	fmt.Println("原始.............")

	println("第十个元素的地址", &data[9])
	println("Main函数中data的地址", &data)

	//给data最后添加元素10
	changeSliceDataCapError(data)

	//Slice数据结构
	//type slice struct {
	//	array unsafe.Pointer
	//  len int
	//  cap int
	//}

	//通过输出可以看出 append 10失败,
	//go只有值传递 。slice在go中是类似引用传递的值传递，相当于New了一个指向data内存对象的slice指针
	//传递slice的时候，这个结构体是值传递的，传递完成后，内存中有两个“slice结构体”，它们引用同一块“slice数组”
	//data[0] = 10  修改这个数组内容，所以里面改了外面能看到改动
	//append这个函数则根据情况不同有两种工作方式：
	// 若你slice的cap足够，则直接修改其引用的数组区域并简单将len增加，并返回这个slice本身，即这时候h=append(h,...)之后，h还是引用原来的数组区域，只不过h的len增加了；
	// 若cap不够，则重新申请一块数组区域并将原来的数组内容拷贝过去后再进行追加元素操作，这时候append的返回的slice引用的是另一块内存了
	fmt.Println("返回.............")
	println("第十个元素的地址", &data[9])
	println("Main函数中data的地址", &data)
	fmt.Println("data:", data)

	//任何时候传指针是可以的
	//https://www.cnblogs.com/snowInPluto/p/7477365.html
	// changeDataByPoint(&data)
	// fmt.Println(data)

	//插入元素
	rear := append([]int{100}, data[5:]...)
	data = append(data[:5], rear...)

	fmt.Println("data:", data)
}
