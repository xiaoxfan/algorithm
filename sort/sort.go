package sort

//外层循环共执行n-1次，每次找出arr[0~n-i]的最大值，并移动到n-i的位置
func BubbleSort(arr []int) []int {
	flag := true
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return arr
}

//判断序列是否正序
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

//基准点在右边
//Lomuto 的分区方案
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[high], arr[i] = arr[i], arr[high]
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}
func QuickSort1(arr []int) []int {
	quickSort1(arr, 0, len(arr)-1)
	return arr
}

//基准点在左边 Hoare 的分区方法
//https://www.jianshu.com/p/116f5b0941f9
func quickSort1(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := arr[low]
	i, j := low-1, high+1
	for {
		for j--; pivot < arr[j]; j-- {
		}
		for i++; pivot > arr[i]; i++ {
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	quickSort1(arr, low, j)
	quickSort1(arr, j+1, high)
}
func QuickSort2(arr []int) []int {
	quickSort2(arr, 0, len(arr)-1)
	return arr
}
//基准点在左边 p记录基准点位置
//从右往左(j--)找到第一个比pivot小的值 arr[p]与arr[j]互换 p=j
//从左往右(i++)找到第一个比pivot大的值 arr[i]与arr[p]互换 p=i
func quickSort2(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := arr[low]
	p, i, j := low, low, high
	for i < j {
		for pivot <= arr[j] && p < j {
			j--
		}
		arr[j], arr[p] = arr[p], arr[j]
		p = j
		for pivot >= arr[i] && i < p {
			i++
		}
		arr[p], arr[i] = arr[i], arr[p]
		p = i
	}
	quickSort2(arr, low, p-1)
	quickSort2(arr, p+1, high)
}

