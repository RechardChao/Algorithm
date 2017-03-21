package main

import (
    "fmt"
    "time"
    "math/rand"
)


func mergeSort(arr []int) []int {
    length := len(arr)
    if length < 2 {
        return arr
    }
    middle := length / 2
    left := arr[0:middle]
    right := arr[middle:]
    return merge(mergeSort(left),mergeSort(right))
}


func merge(left []int,right []int) []int {
    var result []int
    for len(left) != 0 && len(right) != 0 {
        if left[0] < right[0] {
            result = append(result,left[0])
            left = left[1:]

        } else {
            result = append(result,right[0])
            right = right[1:]
        }
    }
    result = append(result,append(left,right...)...)
    return result
}
func main() {
    num := 100
    arr := make([]int,num)
    rand.Seed(int64(time.Now().Nanosecond()))
    for i := range arr {
        arr[i] = rand.Intn(1000)
    }
    fmt.Printf("%v",mergeSort(arr))
    fmt.Println()
    return
}
