package main

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if length > capacity {
		capacity = length
	}
	newslice := make([]int, length, capacity)
	if len(slice) < begin+length {
		length = len(slice) - begin
	}
	end := begin + length
	copy(newslice, slice[begin:end])
	return newslice
}
