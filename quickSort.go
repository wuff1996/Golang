package main

import "fmt"

func quickSort(arr []int, start, end int) {
    if start < end {
        i, j := start, end
        key := arr[start]             //将开始位置记为key
        for i <= j {                  //找到左边大于key的数与右边小于key的数交换
            for arr[i] < key {
                i++
            }
            for arr[j] > key {
                j--
            }
            if i <= j {
                arr[i], arr[j] = arr[j], arr[i]
                i,j = i+1,j-1
            }
        }

        if start < j {               //如果完成排序则递归左边数组
            quickSort(arr, start, j)
        }
        if end > i {                 //如果没完成则递归右边数组
            quickSort(arr, i, end)
        }
    }
}

func main() {

    arr := []int{56565,4554,35,6,454,2,5646,44,5}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}