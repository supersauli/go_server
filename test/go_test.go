// package main
package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

// TestCap
func TestCap(t *testing.T) {
	cap1 := make([]int, 8)
	cap2 := make([]int, 8)
	PrintArrayInfo(&cap1, "cap1")
	PrintArrayInfo(&cap2, "cap2")
	for i := 0; i < len(cap1)-5; i++ {
		cap1[i] = i
	}
	// cap2 会指向cap1的地址
	cap2 = cap1
	PrintArrayInfo(&cap1, "cap1")
	PrintArrayInfo(&cap2, "cap2")

	fmt.Println("cap1=", cap1)
	fmt.Println("cap2=", cap2)
	// 修改cap1，cap2也会修改，因为两个指向同一个地址，所以cap1[0] = 54321不是12345
	cap1[0] = 12345
	cap2[0] = 54321
	fmt.Println("cap1=", cap1)
	fmt.Println("cap2=", cap2)

	PrintArrayInfo(&cap1, "capappp")
	// append 如果超过了最大容量会改变cap1的地址所以cap1和cap2已经不会指向同一个地址了，另外一个例子看TestCap2
	cap1 = append(cap1, 10)
	PrintArrayInfo(&cap1, "capappp")
	cap2 = append(cap2, 11)
	//for i := range cap1 {
	//cap1[i] = i * 10
	//}

	fmt.Println("cap1=", cap1)
	fmt.Println("cap2=", cap2)
	fmt.Println("test copy")

	cap3 := make([]int, 3, 8)
	// cop 并不会修改指向地址
	copy(cap3, cap2)
	cap3[1] = 1111

	PrintArrayInfo(&cap3, "cap3")
	PrintArrayInfo(&cap2, "cap22222")

	fmt.Println("cap3=", cap3)
	fmt.Println("cap2=", cap2)
	cap3 = append(cap3, 22222)
	// cap4和cap3指向同一个地址
	cap4 := cap3[:]
	cap4[0] = 112233
	fmt.Println("cap4=", cap4)
	fmt.Println("cap3=", cap3)

	// 结论切片间用 l=r，l:=r 相当于l和r【数据存储】指向统一个地址，append 如果超过了默认的容量会重新返回一个切片地址，cope(l,r)不会造成 l和r指向同一个地址

}
func TestCap2(t *testing.T) {
	cap1 := make([]int, 4, 8)
	for i := 0; i < len(cap1); i++ {
		cap1[i] = i
	}
	cap2 := cap1

	fmt.Println("*************************************")
	PrintArrayInfo(&cap1, "cap1")
	PrintArrayInfo(&cap2, "cap2")
	// 注意❶
	cap1 = append(cap1, 10)

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	PrintArrayInfo(&cap1, "cap1")
	PrintArrayInfo(&cap2, "cap2")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	// 注意❷
	cap2 = append(cap2, 11)

	PrintArrayInfo(&cap1, "cap1")
	PrintArrayInfo(&cap2, "cap2")

	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")

	// 注意❶ 和注意❷这里有个问题，因为cap1和cap2 数据存储指向统一个位置，但是容量和最大值是各自的值，也就会造成❶append的时候修改了容量是5，cap1容量是5
	//cap2的容量还是4，❷修改的时候因为容量是4，所以它还会修改5的位置，所以结果是11而不是10
	fmt.Println(cap1)
	fmt.Println(cap2)
}
func TestCap3(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5}
	s2 := s1[:1]
	s3 := s1[1:2]
	fmt.Println(s2)
	fmt.Println(s3)
	// 删除 注意s1变了 而且s1也会变成一个[0 2 3 4 5 5]
	//s4 := append(s1[:1], s1[2:]...)
	//fmt.Println(s4)
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1)

	//插入
	s1 = append(s1[0:1], append([]int{1}, s1[1:]...)...)
	fmt.Println(s1)
	// 扩展
	s1 = append(s1, make([]int, 10)...)
	fmt.Println(cap(s1))

}

func PrintArrayInfo(arr *[]int, name string) {
	var p1 = *(*[3]int)(unsafe.Pointer(arr))
	fmt.Println(name, "数据     =", p1[0])
	fmt.Println(name, "已使用   =", p1[1])
	fmt.Println(name, "最大     =", p1[2])
}

func TestArray(t *testing.T) {
	ar1 := make([]int, 5, 8)
	for i := range ar1 {
		ar1[i] = i
	}

	//ar2 := ar1
	var ar2 = make([]int, len(ar1))
	copy(ar2, ar1)

	var p1 = *(*[3]int)(unsafe.Pointer(&ar1))
	var p2 = *(*[3]int)(unsafe.Pointer(&ar2))
	fmt.Println("p1数据		=", p1[0])
	fmt.Println("p1已使用   =", p1[1])
	fmt.Println("p1最大		=", p1[2])

	fmt.Println("p2数据		=", p2[0])
	fmt.Println("p2已使用	=", p2[1])
	fmt.Println("p2最大		=", p2[2])

	ar1 = append(ar1, 100)
	ar2 = append(ar2, 111)

	p1 = *(*[3]int)(unsafe.Pointer(&ar1))
	p2 = *(*[3]int)(unsafe.Pointer(&ar2))

	fmt.Println("p1数据		=", p1[0])
	fmt.Println("p1已使用   =", p1[1])
	fmt.Println("p1最大		=", p1[2])

	fmt.Println("p2数据		=", p2[0])
	fmt.Println("p2已使用   =", p2[1])
	fmt.Println("p2最大		=", p2[2])

	fmt.Println(ar1, len(ar1), cap(ar1))
	fmt.Println(ar2, len(ar2), cap(ar2))

}

func TestArray2(t *testing.T) {
	s1 := [...]int{1, 2, 3}
	s2 := [...]int{4, 5, 6}
	s3 := []int{4, 5, 6}
	if s1 == s2 {
		fmt.Println("xxx")
	}
	var s4 []int
	//var s5 = 0
	//var s6 = []int{nil}
	var s7 = []int{}

	if s4 == nil {
		fmt.Println("s4 nil")
	}

	// if s5 == nil {
	// 	fmt.Println("s5 nil")
	// }

	// if s6 == nil {
	// 	fmt.Println("s6 nil")
	// }

	if s7 == nil {
		fmt.Println("s7 nil")
	}

	// if s1 == s3 {
	// 	error
	// }

	//不同类型nil比较的语义不同。只有slice、map、chan、interface、pointer
	//可以直接和nil比较，array、string、struct不行，但是所有类型都可以转换
	//为接口类型比如interface{}
	// if s1 == nil {
	// 	error
	// }

	if s3 == nil {

	}
	var m2 map[int]int
	if m2 == nil {
		fmt.Println("xxx")
	}

	m1 := make(map[int]int)
	if m1 == nil {
		fmt.Println("xxx")
	}

	// type Iooo struct {
	// }

	//vI := Iooo{}
	// if vI == nil {
	// 	error
	// }
	//var vs string
	// if vs == nil {

	// }

}

func TestSring(t *testing.T) {
	s := "1234567"
	m := []byte(s)
	p1 := *(*[3]int)(unsafe.Pointer(&s))
	p2 := *(*[3]int)(unsafe.Pointer(&m))
	fmt.Println(p1[0])
	fmt.Println(p2[0])
	//s = strings.Replace(s, "1", "9", -1)
	m[0] = '9'
	p1 = *(*[3]int)(unsafe.Pointer(&s))
	p2 = *(*[3]int)(unsafe.Pointer(&m))
	fmt.Println(p1[0])
	fmt.Println(p2[0])

	fmt.Println(s)
	fmt.Println(strconv.Atoi(string(m)))
}
func TestMap(t *testing.T) {
	//s1 := "hellow"
	//fmt.Println(len(s1))
	sd := map[int][]int{}

	if v, ok := sd[1]; ok {
		v = append(v, 2)
	} else {
		sd[1] = []int{11111}
	}

	if v, ok := sd[1]; ok {
		v = append(v, 22222)
		sd[1] = v
	} else {
		sd[1] = []int{23}
	}

	//	sd[12] = append(sd[12], 13)

	for i, j := range sd {
		for _, n := range j {
			fmt.Printf("key=%v,value =%v\n", i, n)
		}
	}

}

type GetInfo interface {
	Out()
}

type BaseData struct {
}

func (d BaseData) Out() {
	fmt.Println("BaseData")
}

type StudentData struct {
}

func (s StudentData) Out() {
	fmt.Println("StudentData")
}

type DataGroup struct {
	Name string
	GetInfo
}

func TestClass(t *testing.T) {
	var d1 = DataGroup{"ba", BaseData{}}
	var d2 = DataGroup{"s", StudentData{}}
	d1.Out()
	d2.Out()

}

func do(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}
func TestChan(t *testing.T) {
	jobs := make(chan int)
	results := make(chan int, 1)
	go func() {
		for v := range jobs {
			time.Sleep(time.Second * 1)
			fmt.Println(v)
		}
		print("xxxxx")
	}()

	for w := 1; w <= 10; w++ {
		jobs <- w
	}
	for w := 1; w <= 3; w++ {
		go do(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	fmt.Printf("close")
	close(jobs)
	fmt.Printf("close")
	for a := 1; a <= 5; a++ {
		v := <-results
		fmt.Println("result", v)
	}
}
