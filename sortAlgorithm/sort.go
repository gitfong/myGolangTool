/*
* sort function for Slice of type int
*/

package sort

//import("fmt")

//冒泡排序，整型列表
//时间复杂度：O(n^2)
//升序
func BubbleSortAsc(values []int) {
	lenth := len(values)
	for i := 0; i < lenth-1; i++ {
		for j := 0; j < lenth-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}
//降序
func BubbleSortDesc(values []int) {
	lenth := len(values)
	for i := 0; i < lenth-1; i++ {
		for j := 0; j < lenth-1-i; j++ {
			if values[j] < values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
}

//快速排序，整型列表
//平均时间复杂度为：O(n * logn)
//T(n) = 2*T(n/2) + n （表示将长度为n的序列划分为两个子序列，每个子序列需要T(n/2)的时间，而划分序列需要n的时间）
//最差时间复杂度为：O(n^2)。 待排序的序列已经有序的情况，快排会下降为冒泡排序；
//升序
func QuickSortAsc(values []int) {
	quickSortAsc(values, 0, len(values)-1)
}
//降序
func QuickSortDesc(values []int){
	quickSortDesc(values,0,len(values)-1)
}
func quickSortAsc(values []int, left int, right int) {
	if left >= right {
		return
	}

	l, r := left, right

	pivot := values[left]
	for left < right {
		for left < right && values[right] >= pivot {
			right--
		}
		values[left], values[right] = values[right], values[left]

		for left < right && values[left] <= pivot {
			left++
		}
		values[left], values[right] = values[right], values[left]
	}
	quickSortAsc(values, l, left-1)
	quickSortAsc(values, left+1, r)
}
func quickSortDesc(values []int, left int, right int) {
	if left >= right {
		return
	}

	l, r := left, right

	pivot := values[left]
	for left < right {
		for left < right && values[right] <= pivot {
			right--
		}
		values[left], values[right] = values[right], values[left]

		for left < right && values[left] >= pivot {
			left++
		}
		values[left], values[right] = values[right], values[left]
	}
	/* 
	//这里不能用goroutine来执行排序任务，否则当前函数返回的时候,排序还未完成。
	go quickSortDesc(values, l, left-1)
	go quickSortDesc(values, left+1, r)
	*/
	quickSortDesc(values, l, left-1)
	quickSortDesc(values, left+1, r)
}

