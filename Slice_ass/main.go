package main

import (
	"fmt"
)

func reverseSlice(slice *[]int) {
	if len(*slice) <= 0 {
		return
	}
	s := *slice
	i := 0
	j := len(*slice) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--

	}
	return

}
func sumValuesByKeys(freq map[int][]int) map[int]int {
	mii := make(map[int]int)
	for key, value := range freq {
		sum := 0
		for _, v := range value {
			sum += v
		}
		mii[key] = sum
	}
	return mii
}
func freqCount(word string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(word); i++ {
		if string(word[i]) != " " {
			m[string(word[i])]++
		}
	}
	return m
}
func main() {
	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	fmt.Println("Original Slice:", slice)
	reverseSlice(&slice)
	fmt.Println("Reversed Slice:", slice)

	// reader:=bufio.NewReader()
	freqMap := freqCount("Nipun Jain")
	for key, value := range freqMap {
		fmt.Printf("Key is:%v\n value is:%v\n", key, value)
	}
	sliceMap := make(map[int][]int)
	sliceMap[1] = []int{1, 2, 3}
	sliceMap[2] = []int{4, 5, 6}
	fmt.Println("Old Map is:", sliceMap)
	newMap := sumValuesByKeys(sliceMap)
	fmt.Println("New Map is:", newMap)

}
