package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	var p1 int
	p1 = 1
	return p1
}

func ReturnFloat() float32 {
	var p2 float32
	p2 = 1.1
	return p2
}

func ReturnIntArray() [3]int {
	p3 := [3]int{1, 3, 4}
	return p3
}

func ReturnIntSlice() []int {
	p4 := []int{1, 2, 3}
	return p4
}

func IntSliceToString(insl []int) string {
	var res = ""
	for _, val := range insl {
		ex := strconv.Itoa(val)
		res = res + ex
	}
	return res
}

func MergeSlices(sl1 []float32, sl2 []int32) []int {
	sl3 := make([]int, 0, len(sl1)+len(sl2))
	for _, val := range sl1 {
		sl3 = append(sl3, int(val))
	}
	for _, val := range sl2 {
		sl3 = append(sl3, int(val))
	}
	return sl3
}

func GetMapValuesSortedByKey(map1 map[int]string) []string {
	var keys []int
	for ind := range map1 {
		keys = append(keys, ind)
	}
	sort.Ints(keys)
	var ressl []string
	for _, k := range keys {
		ressl = append(ressl, map1[k])
	}
	return ressl
}
