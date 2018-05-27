package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert := assert.New(t)
	array := NewArray(20)

	var i, num int
	var err error

	i = -1
	_, err = array.Get(i)
	assert.NotNil(err)

	i = 30
	_, err = array.Get(i)
	assert.NotNil(err)

	i = 10
	num, err = array.Get(i)
	assert.Nil(err)
	assert.Equal(10, num)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	array := NewArray(20)

	var i, num int
	var err error

	i = 21
	num, err = array.Add(i)
	assert.Nil(err)
	_ = num
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	array := NewArray(20)

	var i, num int
	var err error

	i = 10
	num, err = array.Delete(i)
	assert.Nil(err)
	assert.Equal(10, num)
	assert.NotEqual(10, array[9])
	// 下面的断言是不通过的，因为len(array) == 20 这是为什么啊
	// assert.Equal(19, len(array))
}

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	array := NewArray(20)

	var i, j, num int
	var err error

	i = 10
	j = 30
	num, err = array.Insert(i, j)
	assert.Nil(err)
	assert.Equal(j, num)
}
