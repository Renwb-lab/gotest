package main

import "fmt"

func heapSort(arr []int) {
	r := len(arr) - 1
	for i := r / 2; i <= r; i += 1 {
		upAdjust(arr, i, r)
	}
	for i := r; i > 0; i -= 1 {
		arr[i], arr[0] = arr[0], arr[i]
		downAdjust(arr, 0, i-1)
	}
}

func upAdjust(arr []int, cur int, end int) {
	for cur >= 0 {
		if cur+1 <= end && arr[cur+1] > arr[cur] {
			cur += 1
		}
		parentIdx := cur / 2
		if arr[parentIdx] < arr[cur] {
			arr[parentIdx], arr[cur] = arr[cur], arr[parentIdx]
			cur = parentIdx
		} else {
			return
		}
	}
}

func downAdjust(arr []int, cur, end int) {
	for cur <= end {
		leftChildIdx := cur * 2
		if leftChildIdx > end {
			return
		}
		childIdx := leftChildIdx
		rightChildIdx := cur*2 + 1
		if rightChildIdx <= end && arr[rightChildIdx] > arr[leftChildIdx] {
			childIdx = rightChildIdx
		}
		if arr[cur] < arr[childIdx] {
			arr[childIdx], arr[cur] = arr[cur], arr[childIdx]
			cur = childIdx
		} else {
			break
		}
	}
}

func main() {
	arr := []int{5, 6, 2, 2, 3}
	heapSort(arr)
	fmt.Println(arr)
}
