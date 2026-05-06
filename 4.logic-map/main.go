package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(nums, target)
	fmt.Println("Return Index: ", result)

}

func TwoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for currentIndex, val := range nums {
		complement := target - val
		if prevIndex, ok := indexMap[complement]; ok {
			return []int{prevIndex, currentIndex}
		}
		indexMap[val] = currentIndex
	}
	return nil
}
