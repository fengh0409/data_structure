package main

import "fmt"

type Array interface {
	Get(i int) (int, error)
	Add(i int) (int, error)
	Delete(i int) (int, error)
	Insert(i, j int) (int, error)
	Length() int
}

type MyArray []int

func NewArray(maxSize int) MyArray {
	var array = make(MyArray, 20)
	var j int = 1

	for i := 0; i < maxSize; i++ {
		array[i] = j
		j++
	}

	return array
}

func (l MyArray) Get(i int) (int, error) {
	if i < 1 || i > len(l) {
		return 0, fmt.Errorf("the location should greater than zero and less than maxSize")
	}

	return l[i-1], nil
}

func (l MyArray) Add(i int) (int, error) {
	l = append(l, i)

	return len(l), nil
}

func (l MyArray) Delete(i int) (int, error) {
	if i < 1 || i > len(l) {
		return 0, fmt.Errorf("the location should greater than zero and less than maxSize")
	}

	val := l[i-1]
	copy(l[i-1:], l[i:])
	l = l[:len(l)-1]

	return val, nil
}

func (l MyArray) Insert(i, j int) (int, error) {
	if i < 1 || i > len(l) {
		return 0, fmt.Errorf("the location should greater than zero and less than maxSize")
	}

	// 将第i个位置及其后面的内容放入临时切片
	tmpSlice := append(MyArray{}, l[i-1:]...)
	// 第i个位置插入元素
	newSlice := append(l[:i-1], j)
	// 将临时切片追加到新的切片即可
	newSlice = append(newSlice, tmpSlice...)

	return newSlice[i-1], nil
}

func (l MyArray) Length() int {
	return len(l)
}
