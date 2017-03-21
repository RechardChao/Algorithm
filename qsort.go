package main

import (
    "fmt"
)

func main() {
    arr := []int{9,8,3,4,5,6,6,6,63,34,56,4,3,2,1,13,4,5,5,5,9,9,99,56,18,1}
    QuickSort(arr)
    fmt.Printf("%v",arr)
    fmt.Println()
}

func  QuickSort(arr []int) {
    quickSort(arr,0,len(arr)-1)
}

func quickSort(array []int, left int, right int) {
  if left < right {
    key := array[left]
    low := left
    high := right
    for low < high {
        for low < high   array[high] >= key {
            
        high--
      }
      array[low] = array[high]
      for  array[low] <= key {
        low++
      }
      array[high] = array[low]
    }
    array[low] = key
    quickSort(array, left, low-1);
    quickSort(array, low+1, right);
  }
}
