package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeOf_Basic(t *testing.T) {
	var (
		x int64 = 123
		y int8  = 1
	)

	fmt.Printf("int64 pointer size: %d\n", unsafe.Sizeof(&x))
	fmt.Printf("int64 value size: %d\n", unsafe.Sizeof(x))

	fmt.Printf("int8 pointer size: %d\n", unsafe.Sizeof(&y))
	fmt.Printf("int8 value size: %d\n", unsafe.Sizeof(y))

	// type stringHeader struct {
	// 		Data uintptr
	// 		Len  int
	// }
	var (
		ss = "ab"
		bs = "abcdabcdabcdabcdabcdabcdabcdabcdsabcdabcdabcsdfdsd"
	)

	fmt.Printf("string pointer size: %d\n", unsafe.Sizeof(&ss))
	fmt.Printf("string value size: %d\n", unsafe.Sizeof(ss))
	fmt.Printf("string len: %d\n", len(ss))

	fmt.Printf("string pointer size: %d\n", unsafe.Sizeof(&bs))
	fmt.Printf("string value size: %d\n", unsafe.Sizeof(bs))
	fmt.Printf("string len: %d\n", len(bs))

	// type sliceHeader struct {
	// 	array unsafe.Pointer // 元素指標
	// 	len   int            // 長度
	// 	cap   int            // 容量
	// }
	var (
		intSlice = []int64{1, 2, 3, 4, 6}
		strSlice = []string{"1", "2", "12345", "3"}
	)

	fmt.Printf("int slice pointer size: %d\n", unsafe.Sizeof(&intSlice))
	fmt.Printf("int slice value size: %d\n", unsafe.Sizeof(intSlice))
	fmt.Printf("int slice len: %d\n", len(intSlice))

	fmt.Printf("string slice pointer size: %d\n", unsafe.Sizeof(&strSlice))
	fmt.Printf("string slice value size: %d\n", unsafe.Sizeof(strSlice))
	fmt.Printf("string slice len: %d\n", len(strSlice))

	var (
		strMap = map[string]string{"a": "b"}
	)

	fmt.Printf("string map pointer size: %d\n", unsafe.Sizeof(&strMap))
	fmt.Printf("string map value size: %d\n", unsafe.Sizeof(strMap))
	fmt.Printf("string map len: %d\n", len(strMap))
}

func TestSizeOf_IterateArr(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	itemSize := unsafe.Sizeof(int32(0))

	for i := 0; i < len(arr); i++ {
		fmt.Printf("item %d is %d\n", i,
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(i)*uintptr(itemSize))))
	}
}

func TestOffset_ChangeField(t *testing.T) {
	type intsStruct struct {
		a int
		b int
		c int
	}

	is := intsStruct{
		a: 1, b: 2, c: 3,
	}
	aPtr := unsafe.Pointer(&is.a)
	beforeB := is.b

	bPtr := (*int)(unsafe.Pointer(uintptr(aPtr) + unsafe.Offsetof(is.b)))
	*bPtr = 1
	fmt.Printf("struct's b-field: from %d to %d\n", beforeB, is.b)
}

func TestAlignOf_Demo(t *testing.T) {
	type intsStruct struct {
		a int
		b string
		c int
	}
	is := intsStruct{
		a: 1, b: "2", c: 3,
	}
	t.Log(unsafe.Alignof(is.a))
}
