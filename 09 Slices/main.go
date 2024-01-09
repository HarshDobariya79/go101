package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{}

	nums = append(nums, 1, 5, 3)

	fmt.Println(nums)

	nums = append(nums[1:], 6, 5)

	fmt.Println(nums)

	var items = make([]int, 4)

	fmt.Println(items)

	items = append(items, 3, 6, 3)

	fmt.Println(items)

	fmt.Println(sort.IntsAreSorted(items))

	sort.Ints(items)

	fmt.Println(items)

	fmt.Println(sort.IntsAreSorted(items))

	index := 4

	items = append(items[:index], items[index+1:]...)

	fmt.Println(items)
}
