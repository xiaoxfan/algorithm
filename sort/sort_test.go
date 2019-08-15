package sort

import (
	"fmt"
	"testing"
)
var arr=[]int{1,9,8,2,7,4,6,5,9,1,8,2,7,3,6,4,5,1,9,8,2,7,4,6,5,9,1,8,2,7,3,6,4,5,1,9,8,2,7,4,6,5,9,1,8,2,7,3,6,4,5}
func TestBubbleSort(t *testing.T) {
	ret := BubbleSort(arr)
	fmt.Println(IsSorted(ret))
}
func TestIsSorted(t *testing.T) {
	fmt.Println(IsSorted(arr))
	fmt.Println(IsSorted(BubbleSort(arr)))
}
func TestQuickSort(t *testing.T) {
	ret := QuickSort(arr)
	fmt.Println(IsSorted(ret))
}
func TestQuickSort1(t *testing.T) {
	ret := QuickSort1(arr)
	fmt.Println(IsSorted(ret))
}
func TestQuickSort2(t *testing.T) {
	ret := QuickSort2(arr)
	fmt.Println(IsSorted(ret))
}