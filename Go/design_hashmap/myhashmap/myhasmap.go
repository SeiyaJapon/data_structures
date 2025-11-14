package myhashmap

import (
	"strconv"
	"strings"
)

type MyHashMap struct {
	myMap [][]int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	if this.Get(key) != -1 {
		for _, item := range this.myMap {
			if item[0] == key {
				item[1] = value
			}
		}
		return
	}

	this.myMap = append(this.myMap, []int{key, value})
}

func (this *MyHashMap) Get(key int) int {
	for _, item := range this.myMap {
		if item[0] == key {
			return item[1]
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	for i, item := range this.myMap {
		if item[0] == key {
			this.myMap = append(this.myMap[:i], this.myMap[i+1:]...)
			return
		}
	}
}

func (m MyHashMap) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i, item := range m.myMap {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString("[")
		for j, v := range item {
			if j > 0 {
				b.WriteString(",")
			}
			b.WriteString(strconv.Itoa(v))
		}
		b.WriteString("]")
	}
	b.WriteString("]")
	return b.String()
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
